package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vasapolrittideah/accord/apps/api/features/auth/usecase"
	"github.com/vasapolrittideah/accord/apps/api/internal/config"
	"github.com/vasapolrittideah/accord/apps/api/internal/response"
	validate "github.com/vasapolrittideah/accord/apps/api/internal/validator"
)

type AuthHandler struct {
	usecase usecase.AuthUseCase
	conf    *config.Config
}

func NewAuthHandler(service usecase.AuthUseCase, conf *config.Config) AuthHandler {
	return AuthHandler{service, conf}
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	payload := new(usecase.SignUpRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error("The request is in a invalid format"))
	}

	if errs := validate.ValidateStruct(payload, h.conf.ValidationTranslator); len(errs) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail(errs))
	}

	user, err := h.usecase.SignUp(*payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(user))
}

func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	payload := new(usecase.SignInRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error("The request is in a invalid format"))
	}

	if errs := validate.ValidateStruct(payload, h.conf.ValidationTranslator); len(errs) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.Fail(errs))
	}

	token, err := h.usecase.SignIn(*payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(token))
}

func (h AuthHandler) SignOut(c *fiber.Ctx) error {
	uuidString := c.Locals("sub")
	userId, _ := uuid.Parse(uuidString.(string))

	user, err := h.usecase.SignOut(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(user))
}

func (h AuthHandler) RefreshToken(c *fiber.Ctx) error {
	uuidString := c.Locals("sub")
	userId, _ := uuid.Parse(uuidString.(string))
	refreshToken := c.Locals("refresh_token").(string)

	// Generate new access token and refresh token
	token, err := h.usecase.RefreshToken(userId, refreshToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(token))
}
