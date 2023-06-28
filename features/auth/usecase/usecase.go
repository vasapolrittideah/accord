package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vasapolrittideah/accord/features/auth/repository"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/utils"
	"github.com/vasapolrittideah/accord/models"
	"strings"
)

//go:generate mockery --name AuthUseCase --filename usecase_mock.go
type AuthUseCase interface {
	SignUp(payload SignUpRequest) (*models.User, error)
	SignIn(payload SignInRequest) (*Tokens, error)
	SignOut(userId uuid.UUID) (*models.User, error)
	RefreshToken(userId uuid.UUID, userRefreshToken string) (*Tokens, error)
	ParseToken(tokenString, tokenPublicKey string) (*jwt.MapClaims, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
	conf     *config.Config
}

type SignUpRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=32"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8,max=32"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthUseCase(repo repository.UserRepository, conf *config.Config) AuthUseCase {
	return authUseCase{repo, conf}
}

func (u authUseCase) SignUp(payload SignUpRequest) (*models.User, error) {
	if payload.Password != payload.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, errors.New("unable to hash the user password")
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     "USER",
		Verified: false,
		Provider: "local",
	}

	user, err := u.userRepo.CreateUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email does already exist")
		}

		return nil, errors.New("unable to create a new user")
	}

	return &user, nil
}

func (u authUseCase) SignIn(payload SignInRequest) (*Tokens, error) {
	user, err := u.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return nil, errors.New("email does not exist")
	}

	if ok, err := utils.VerifyPassword(user.Password, payload.Password); err != nil || !ok {
		return nil, errors.New("password is not correct")
	}

	accessToken, err := utils.GenerateToken(
		u.conf.AccessTokenExpiresIn,
		u.conf.AccessTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	refreshToken, err := utils.GenerateToken(
		u.conf.RefreshTokenExpiresIn,
		u.conf.RefreshTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	hashedRefreshToken, err := utils.HashRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("unable to hash newly generated token")
	}

	if _, err = u.userRepo.UpdateUser(user.ID, models.User{RefreshToken: hashedRefreshToken}); err != nil {
		return nil, errors.New("unable to store token in the database")
	}

	tokens := &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (u authUseCase) SignOut(userId uuid.UUID) (*models.User, error) {
	user, err := u.userRepo.UpdateUser(userId, models.User{RefreshToken: ""})
	if err != nil {
		return nil, errors.New("unable to update user to the database")
	}

	return &user, nil
}

func (u authUseCase) RefreshToken(userId uuid.UUID, userRefreshToken string) (*Tokens, error) {
	user, err := u.userRepo.GetUser(userId)
	if err != nil {
		return nil, errors.New("unable to get user from the database")
	}

	if ok, err := utils.VerifyRefreshToken(user.RefreshToken, userRefreshToken); err != nil || !ok {
		return nil, errors.New("verifying refresh token has been failed")
	}

	accessToken, err := utils.GenerateToken(
		u.conf.AccessTokenExpiresIn,
		u.conf.AccessTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	refreshToken, err := utils.GenerateToken(
		u.conf.RefreshTokenExpiresIn,
		u.conf.RefreshTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	hashedRefreshToken, _ := utils.HashRefreshToken(refreshToken)
	_, err = u.userRepo.UpdateUser(userId, models.User{RefreshToken: hashedRefreshToken})
	if err != nil {
		return nil, errors.New("unable to update user to the database")
	}

	tokens := &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (u authUseCase) ParseToken(tokenString, tokenPublicKey string) (*jwt.MapClaims, error) {
	token, err := utils.ValidateToken(tokenString, tokenPublicKey)
	if err != nil {
		return nil, errors.New("token is invalid or has been expired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return &claims, nil
}
