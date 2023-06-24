package service

import (
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
	"github.com/vasapolrittideah/accord/features/auth/repository"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/utils"
	"github.com/vasapolrittideah/accord/models"
	"strings"
	"time"
)

//go:generate mockery --name AuthService --filename service_mock.go
type AuthService interface {
	SignUp(payload SignUpRequest) (*models.User, error)
	SignIn(payload SignInRequest) (*Tokens, error)
	SignOut(userId uuid.UUID) (*models.User, error)
	RefreshToken(userId uuid.UUID, userRefreshToken string) (*Tokens, error)
	GenerateToken(ttl time.Duration, privateKey string, userId uuid.UUID) (string, error)
	ValidateToken(token string, publicKey string) (*jwt.Token, error)
	HashRefreshToken(refreshToken string) (string, error)
	VerifyRefreshToken(encoded string, refreshToken string) (bool, error)
}

type authService struct {
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

func NewAuthService(repo repository.UserRepository, conf *config.Config) AuthService {
	return authService{repo, conf}
}

func (s authService) SignUp(payload SignUpRequest) (*models.User, error) {
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

	user, err := s.userRepo.CreateUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email does already exist")
		}

		return nil, errors.New("unable to create a new user")
	}

	return &user, nil
}

func (s authService) SignIn(payload SignInRequest) (*Tokens, error) {
	user, err := s.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return nil, errors.New("email does not exist")
	}

	if ok, err := utils.VerifyPassword(user.Password, payload.Password); err != nil || !ok {
		return nil, errors.New("password is not correct")
	}

	accessToken, err := s.GenerateToken(
		s.conf.AccessTokenExpiresIn,
		s.conf.AccessTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	refreshToken, err := s.GenerateToken(
		s.conf.RefreshTokenExpiresIn,
		s.conf.RefreshTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	hashedRefreshToken, err := s.HashRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("unable to hash newly generated token")
	}

	if _, err = s.userRepo.UpdateUser(user.ID, models.User{RefreshToken: hashedRefreshToken}); err != nil {
		return nil, errors.New("unable to store token in the database")
	}

	tokens := &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (s authService) SignOut(userId uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.UpdateUser(userId, models.User{RefreshToken: ""})
	if err != nil {
		return nil, errors.New("unable to update user to the database")
	}

	return &user, nil
}

func (s authService) RefreshToken(userId uuid.UUID, userRefreshToken string) (*Tokens, error) {
	user, err := s.userRepo.GetUser(userId)
	if err != nil {
		return nil, errors.New("unable to get user from the database")
	}

	if ok, err := s.VerifyRefreshToken(user.RefreshToken, userRefreshToken); err != nil || !ok {
		return nil, errors.New("verifying refresh token has been failed")
	}

	accessToken, err := s.GenerateToken(
		s.conf.AccessTokenExpiresIn,
		s.conf.AccessTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	refreshToken, err := s.GenerateToken(
		s.conf.RefreshTokenExpiresIn,
		s.conf.RefreshTokenPrivateKey,
		user.ID,
	)
	if err != nil {
		return nil, errors.New("failed to generate jwt token: " + err.Error())
	}

	hashedRefreshToken, _ := s.HashRefreshToken(refreshToken)
	_, err = s.userRepo.UpdateUser(userId, models.User{RefreshToken: hashedRefreshToken})
	if err != nil {
		return nil, errors.New("unable to update user to the database")
	}

	tokens := &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (s authService) GenerateToken(ttl time.Duration, privateKey string, userId uuid.UUID) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", errors.New("unable to decode key: " + err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decoded)
	if err != nil {
		return "", errors.New("unable to parse key: " + err.Error())
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userId,
		"iat": now.Unix(),
		"exp": now.Add(ttl).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", errors.New("unable to sign jwt token: " + err.Error())
	}

	return token, nil
}

func (s authService) ValidateToken(token string, publicKey string) (*jwt.Token, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, errors.New("unable to decode key: " + err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, errors.New("unable to parse key: " + err.Error())
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method: " + t.Header["alg"].(string))
		}
		return key, nil
	})
}

func (s authService) HashRefreshToken(refreshToken string) (string, error) {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(refreshToken))
	if err != nil {
		return "", errors.New("unable to hash refresh token: " + err.Error())
	}

	return string(encoded), nil
}

func (s authService) VerifyRefreshToken(encoded string, refreshToken string) (bool, error) {
	return argon2.VerifyEncoded([]byte(refreshToken), []byte(encoded))
}
