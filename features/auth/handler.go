package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vasapolrittideah/accord/internal/config"
	"github.com/vasapolrittideah/accord/internal/response"
	validate "github.com/vasapolrittideah/accord/internal/validator"
)

func RegisterHandlers(r fiber.Router, service Service, conf *config.Config) {
	handler := newHandler(service, conf)
	router := r.Group("/auth")

	router.Post("/signup", handler.signUp)
}

type handler struct {
	service Service
	conf    *config.Config
}

func newHandler(service Service, conf *config.Config) handler {
	return handler{service, conf}
}

func (h handler) signUp(c *fiber.Ctx) error {
	payload := new(SignUpRequest)

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
