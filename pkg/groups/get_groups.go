package groups

import (
	"fiber/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetGrupos(app *fiber.Ctx) error {

	var grupos []models.Groups
	err := r.Db.Find(&grupos).Error

	if err != nil {
		app.Status(404).JSON(&fiber.Map{
			"message": "Grupos not found",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(grupos),
		"next":     "null",
		"previous": "null",
		"items":    grupos,
	})
}
