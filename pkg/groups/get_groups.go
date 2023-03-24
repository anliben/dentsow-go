package groups

import (
	"github.com/gofiber/fiber/v2"

)

func (r handler) GetGrupos(app *fiber.Ctx) error {

	return app.JSON(&fiber.Map{
		"count":    0,
		"next":     "null",
		"previous": "null",
		// "items":    groups,
	})
}
