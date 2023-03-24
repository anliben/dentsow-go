package procedure

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r handler) GetById(app *fiber.Ctx) error {
	var procedure models.Procedure
	id := app.Params("id")

	if id == "" {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID is empty",
		})
		return nil
	}

	err := r.Db.First(&procedure).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "User not found",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"item": procedure,
	})
}
