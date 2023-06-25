package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/features/auth/usecase"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/response"
	"strings"
)

//go:generate mockery --name AuthMiddleware --filename middleware_mock.go
type AuthMiddleware interface {
	AuthenticateWithAccessToken(conf *config.Config) fiber.Handler
}

type authMiddleware struct {
	usecase usecase.AuthUseCase
}

func NewAuthMiddleware(usecase usecase.AuthUseCase) AuthMiddleware {
	return authMiddleware{usecase}
}

func (m authMiddleware) AuthenticateWithAccessToken(conf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const Bearer = "Bearer"
		authHandler := c.Get("Authorization")
		if authHandler == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("No Authorization header found"))
		}

		headerParts := strings.Split(authHandler, " ")
		if len(headerParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("Malformed Authorization header"))
		}

		if headerParts[0] != Bearer {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("Malformed Authorization header"))
		}

		claims, err := m.usecase.ParseToken(headerParts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error(err.Error()))
		}

		c.Locals("sub", (*claims)["sub"])
		return c.Next()
	}
}
