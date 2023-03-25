package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var budget []models.Budget
	err := r.Db.Preload("Cliente").Preload("Vendedor").Preload("Arquivos").Preload("Procedure").Preload("ValorProposta").Find(&budget).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(budget),
		"next":     "null",
		"previous": "null",
		"items":    budget,
	})
}
