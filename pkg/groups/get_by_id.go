package groups

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetById(app *fiber.Ctx) error {
	var grupos models.Groups

	err := r.Db.Find(&grupos).Where("id = ?", app.Params("id")).First(&grupos).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    1,
		"next":     "null",
		"previous": "null",
		"items":    grupos,
	})
}
