package utils

import (
	"github.com/gofiber/fiber/v2"
)

func (r handler) Migrate(app *fiber.Ctx) error {

	return app.JSON(&fiber.Map{
		"detail": "Migrate success",
	})
}
