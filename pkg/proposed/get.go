package proposed

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *handler) GetAll(app *fiber.Ctx) error {
	var proposed []models.ProposedValue

	err := r.Db.Find(&proposed).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Proposed not found",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(proposed),
		"next":     "null",
		"previous": "null",
		"items":    proposed,
	})
}
