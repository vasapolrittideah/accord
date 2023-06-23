package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/features/auth/service"
	"github.com/vasapolrittideah/accord/internal/config"
)

func RegisterHandlers(r fiber.Router, service service.AuthService, conf *config.Config) {
	authHandler := NewAuthHandler(service, conf)
	router := r.Group("/auth")

	router.Post("/signup", authHandler.SignUp)
	router.Post("/signin", authHandler.SignIn)
}
