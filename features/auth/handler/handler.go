package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/features/auth/service"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/response"
	validate "github.com/vasapolrittideah/accord/internal/validator"
)

type AuthHandler struct {
	service service.AuthService
	conf    *config.Config
}

func NewAuthHandler(service service.AuthService, conf *config.Config) AuthHandler {
	return AuthHandler{service, conf}
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	payload := new(service.SignUpRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error("The request is in a invalid format"))
	}

	if errs := validate.ValidateStruct(payload, h.conf.ValidationTranslator); len(errs) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail(errs))
	}

	user, err := h.service.SignUp(*payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(user))
}

func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	payload := new(service.SignInRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error("The request is in a invalid format"))
	}

	if errs := validate.ValidateStruct(payload, h.conf.ValidationTranslator); len(errs) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail(errs))
	}

	token, err := h.service.SignIn(*payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(token))
}
