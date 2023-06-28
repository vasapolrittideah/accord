package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/features/auth/middleware"
	"github.com/vasapolrittideah/accord/features/auth/usecase"
	"github.com/vasapolrittideah/accord/internal/config"
)

func RegisterHandlers(r fiber.Router, conf *config.Config, service usecase.AuthUseCase, m middleware.AuthMiddleware) {
	authHandler := NewAuthHandler(service, conf)
	router := r.Group("/auth")

	router.Post("/signup", authHandler.SignUp)
	router.Post("/signin", authHandler.SignIn)
	router.Post("/signout", m.AuthenticateWithJWT(conf, middleware.Access), authHandler.SignOut)
	router.Post("/refresh", m.AuthenticateWithJWT(conf, middleware.Refresh), authHandler.RefreshToken)
}
