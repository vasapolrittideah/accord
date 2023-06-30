package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/usecase"
	"github.com/vasapolrittideah/accord/apps/api/internal/config"
	"github.com/vasapolrittideah/accord/apps/api/internal/response"
	"strings"
)

//go:generate mockery --name AuthMiddleware --filename middleware_mock.go
type AuthMiddleware interface {
	AuthenticateWithJWT(conf *config.Config, tokenType TokenType) fiber.Handler
}

type authMiddleware struct {
	usecase usecase.AuthUseCase
}

func NewAuthMiddleware(usecase usecase.AuthUseCase) AuthMiddleware {
	return authMiddleware{usecase}
}

type TokenType int

const (
	Access TokenType = iota
	Refresh
)

func (m authMiddleware) AuthenticateWithJWT(conf *config.Config, tokenType TokenType) fiber.Handler {
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

		var publicKey string
		if tokenType == Access {
			publicKey = conf.AccessTokenPublicKey
		} else if tokenType == Refresh {
			publicKey = conf.RefreshTokenPublicKey
		}

		claims, err := m.usecase.ParseToken(headerParts[1], publicKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error(err.Error()))
		}

		c.Locals("token", headerParts[1])
		c.Locals("sub", (*claims)["sub"])
		return c.Next()
	}
}
