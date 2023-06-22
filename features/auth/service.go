package auth

import (
	"errors"
	"github.com/vasapolrittideah/accord/internal/utils"
	"github.com/vasapolrittideah/accord/models"
	"strings"
)

type Service interface {
	SignUp(payload SignUpRequest) (*models.User, error)
}

type service struct {
	repo Repository
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

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) SignUp(payload SignUpRequest) (*models.User, error) {
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

	user, err := s.repo.CreateUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("email does already exist")
		}

		return nil, errors.New("unable to create a new user")
	}

	return &user, nil
}
