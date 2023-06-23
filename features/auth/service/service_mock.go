package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vasapolrittideah/accord/models"
	"time"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) SignUp(payload SignUpRequest) (*models.User, error) {
	args := m.Called(payload)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *AuthServiceMock) SignIn(payload SignInRequest) (*Tokens, error) {
	args := m.Called(payload)
	return args.Get(0).(*Tokens), args.Error(1)
}

func (m *AuthServiceMock) GenerateToken(ttl time.Duration, privateKey string, userId uuid.UUID) (string, error) {
	args := m.Called(ttl, privateKey, userId)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) ValidateToken(token string, publicKey string) (*jwt.Token, error) {
	args := m.Called(token, publicKey)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

func (m *AuthServiceMock) HashRefreshToken(refreshToken string) (string, error) {
	args := m.Called(refreshToken)
	return args.String(0), args.Error(1)
}

func (m *AuthServiceMock) VerifyRefreshToken(encoded string, refreshToken string) (bool, error) {
	args := m.Called(encoded, refreshToken)
	return args.Bool(0), args.Error(1)
}
