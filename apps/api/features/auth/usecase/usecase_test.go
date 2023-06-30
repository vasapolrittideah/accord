package usecase

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/repository"
	"github.com/vasapolrittideah/accord/apps/api/internal/config"
	"github.com/vasapolrittideah/accord/apps/api/models"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := os.Setenv("ENVIRONMENT", "test"); err != nil {
		log.Fatalln(err)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestAuthFlow(t *testing.T) {
	userRepo := repository.NewMockUserRepository(t)

	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	authUseCase := NewAuthUseCase(userRepo, conf)

	// A user that will be returned from the database
	mockUser := models.User{
		ID:        uuid.New(),
		Name:      "TestUser",
		Email:     "testuser@admin.com",
		Role:      "USER",
		Provider:  "local",
		Verified:  false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	signUpBody := SignUpRequest{
		Name:            mockUser.Name,
		Email:           mockUser.Email,
		Password:        "P@ssword!",
		PasswordConfirm: "P@ssword!",
	}

	// First, sign up a new user. Password will be hashed and stored in the database.
	userRepo.EXPECT().CreateUser(mock.AnythingOfType("User")).Return(mockUser, nil).Run(func(_user models.User) {
		mockUser.Password = _user.Password
	})
	_, err = authUseCase.SignUp(signUpBody)
	assert.NoError(t, err)

	signInBody := SignInRequest{
		Email:    signUpBody.Email,
		Password: signUpBody.Password,
	}

	// Second, the user signs in. Refresh token will be generated and stored in the database.
	userRepo.EXPECT().GetByEmail(mock.AnythingOfType("string")).Return(mockUser, nil)
	userRepo.EXPECT().UpdateUser(mock.AnythingOfType("UUID"), mock.AnythingOfType("User")).
		Return(mockUser, nil).Run(func(id uuid.UUID, newData models.User) {
		mockUser.RefreshToken = newData.RefreshToken
	})
	tokens, err := authUseCase.SignIn(signInBody)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)

	// Third, the user's refresh token will be refreshed when it's time.
	userRepo.EXPECT().GetUser(mock.Anything).Return(mockUser, nil)
	newTokens, err := authUseCase.RefreshToken(mockUser.ID, tokens.RefreshToken)
	assert.NoError(t, err)
	assert.NotEqual(t, newTokens.RefreshToken, tokens.RefreshToken)

	// Finally, the user logs out. Refresh token will be emptied in the database.
	_, err = authUseCase.SignOut(mockUser.ID)
	assert.Equal(t, "", mockUser.RefreshToken)
}
