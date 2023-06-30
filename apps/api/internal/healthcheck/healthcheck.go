package healthcheck

import "github.com/gofiber/fiber/v2"

func RegisterHandler(app *fiber.App, version string) {
	app.Get("/healthcheck", healthCheck(version))
}

func healthCheck(version string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("OK " + version)
	}
}
