package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vasapolrittideah/accord/features/auth/usecase"
	"github.com/vasapolrittideah/accord/internal/config"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := os.Setenv("ENVIRONMENT", "test"); err != nil {
		log.Fatalln(err)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestAuthMiddleware_AuthenticateWithAccessToken(t *testing.T) {
	app := fiber.New()
	mockAuthUsecase := usecase.NewMockAuthUseCase(t)

	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	app.Post(
		"/api/endpoint",
		NewAuthMiddleware(mockAuthUsecase).AuthenticateWithJWT(conf, Access),
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).SendString("OK")
		},
	)

	var (
		method = "POST"
		target = "/api/endpoint"
	)

	// No Authentication header request
	req := httptest.NewRequest(method, target, nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	// Empty Authentication header request
	req = httptest.NewRequest(method, target, nil)
	req.Header.Set("Authorization", "")
	resp, _ = app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	// Bearer Authentication header with no token request
	mockAuthUsecase.EXPECT().ParseToken("-", mock.AnythingOfType("string")).Return(nil, errors.New("token is invalid or has been expired"))
	req = httptest.NewRequest(method, target, nil)
	req.Header.Set("Authorization", "Bearer -")
	resp, _ = app.Test(req)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	// Valid Authentication header request
	mockAuthUsecase.EXPECT().ParseToken("token", mock.AnythingOfType("string")).Return(&jwt.MapClaims{}, nil)
	req = httptest.NewRequest(method, target, nil)
	req.Header.Set("Authorization", "Bearer token")
	resp, _ = app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
