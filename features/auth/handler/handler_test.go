package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vasapolrittideah/accord/features/auth/middleware"
	"github.com/vasapolrittideah/accord/features/auth/usecase"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/test"
	"github.com/vasapolrittideah/accord/models"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthHandler_SignUn(t *testing.T) {
	app := fiber.New()
	mockAuthUsecase := usecase.NewMockAuthUseCase(t)
	mockAuthMiddleware := middleware.NewMockAuthMiddleware(t)

	mockAuthMiddleware.EXPECT().AuthenticateWithAccessToken(&config.Config{}).Return(
		func(c *fiber.Ctx) error {
			return c.Next()
		},
	)

	RegisterHandlers(app, &config.Config{}, mockAuthUsecase, mockAuthMiddleware)

	signUpBody := usecase.SignUpRequest{
		Name:            "test",
		Email:           "test@admin.com",
		Password:        "P@ssword!",
		PasswordConfirm: "P@ssword!",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	user := &models.User{
		ID:        uuid.New(),
		Name:      signUpBody.Name,
		Email:     signUpBody.Email,
		Role:      "USER",
		Provider:  "local",
		Verified:  false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	mockAuthUsecase.EXPECT().SignUp(signUpBody).Return(user, nil)

	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	data, _ := test.GetDataFromResponse[models.User](resp)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, user, data)
}

func TestAuthHandler_SignIn(t *testing.T) {
	app := fiber.New()
	mockAuthUsecase := usecase.NewMockAuthUseCase(t)
	mockAuthMiddleware := middleware.NewMockAuthMiddleware(t)

	mockAuthMiddleware.EXPECT().AuthenticateWithAccessToken(&config.Config{}).Return(
		func(c *fiber.Ctx) error {
			return c.Next()
		},
	)

	RegisterHandlers(app, &config.Config{}, mockAuthUsecase, mockAuthMiddleware)

	signInBody := usecase.SignInRequest{
		Email:    "test@admin.com",
		Password: "P@ssword!",
	}

	body, err := json.Marshal(signInBody)
	assert.NoError(t, err)

	token := &usecase.Tokens{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
	}

	mockAuthUsecase.EXPECT().SignIn(signInBody).Return(token, nil)

	req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	data, _ := test.GetDataFromResponse[usecase.Tokens](resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, token, data)
}

func TestAuthHandler_SignOut(t *testing.T) {
	app := fiber.New()
	authService := usecase.NewMockAuthUseCase(t)
	authMiddleware := middleware.NewMockAuthMiddleware(t)

	userId := uuid.New()

	// Assume that the user is authorized
	authMiddleware.On("AuthenticateWithAccessToken", &config.Config{}).Return(
		func(c *fiber.Ctx) error {
			c.Locals("sub", userId.String())
			return c.Next()
		},
	)

	RegisterHandlers(app, &config.Config{}, authService, authMiddleware)

	user := &models.User{
		ID:        uuid.New(),
		Name:      "test",
		Email:     "test@admin.com",
		Role:      "USER",
		Provider:  "local",
		Verified:  false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	authService.On("SignOut", userId).Return(user, nil)

	req := httptest.NewRequest("POST", "/auth/signout", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	data, _ := test.GetDataFromResponse[models.User](resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, user, data)
}
