package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vasapolrittideah/accord/features/auth/service"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/test"
	"github.com/vasapolrittideah/accord/models"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSignUp(t *testing.T) {
	app := fiber.New()
	authService := new(service.AuthServiceMock)

	RegisterHandlers(app, authService, &config.Config{})

	signUpBody := service.SignUpRequest{
		Name:            "test",
		Email:           "test@admin.com",
		Password:        "P@ssword!",
		PasswordConfirm: "P@ssword!",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	user := &models.User{
		ID:        uuid.New(),
		Name:      "Kim",
		Email:     "kim@gmail.com",
		Role:      "USER",
		Provider:  "local",
		Verified:  false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	authService.On("SignUp", signUpBody).Return(user, nil)

	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	data, err := test.GetDataFromResponse[models.User](resp)

	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.Equal(t, user, data)
}

func TestSignIn(t *testing.T) {
	app := fiber.New()
	authService := new(service.AuthServiceMock)

	RegisterHandlers(app, authService, &config.Config{})

	signInBody := service.SignInRequest{
		Email:    "test@admin.com",
		Password: "P@ssword!",
	}

	body, err := json.Marshal(signInBody)
	assert.NoError(t, err)

	token := &service.Tokens{
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
	}

	authService.On("SignIn", signInBody).Return(token, nil)

	req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	data, err := test.GetDataFromResponse[service.Tokens](resp)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, token, data)
}
