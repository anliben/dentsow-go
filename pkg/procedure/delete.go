package procedure

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Delete(app *fiber.Ctx) error {
	var procedimento models.Procedure

	id := app.Params("id")

	if id == "" {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID is empty",
		})
		return nil
	}

	err := r.Db.Where("id = ?", id).Delete(&procedimento).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Procedimento not found",
		})
		return err
	}
	return app.Status(http.StatusNoContent).JSON(&fiber.Map{})
}
