package auth

import (
	"github.com/stretchr/testify/mock"
	"github.com/vasapolrittideah/accord/models"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) SignUp(payload SignUpRequest) (*models.User, error) {
	args := m.Called(payload)

	return args.Get(0).(*models.User), args.Error(1)
}
