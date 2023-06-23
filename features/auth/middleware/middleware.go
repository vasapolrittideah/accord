package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vasapolrittideah/accord/features/auth/service"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/response"
)

type AuthMiddleware interface {
	AuthenticateWithAccessToken(conf *config.Config) fiber.Handler
}

type authMiddleware struct {
	service service.AuthService
}

func NewAuthMiddleware(service service.AuthService) AuthMiddleware {
	return authMiddleware{service}
}

func (m authMiddleware) AuthenticateWithAccessToken(conf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		const Bearer string = "Bearer "
		authHandler := c.Get("Authorization")
		if authHandler == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("No Authorization header found"))
		}

		tokenString := authHandler[len(Bearer):]
		token, err := m.service.ValidateToken(tokenString, conf.AccessTokenPublicKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("Token is invalid or has been expired"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(
				response.Error("Token is invalid"))
		}

		c.Locals("sub", claims["sub"])
		return c.Next()
	}
}
