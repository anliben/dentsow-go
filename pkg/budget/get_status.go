package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetStatus(app *fiber.Ctx) error {
	var orcamento models.Budget

	id := app.Params("id")
	status := app.Params("id")

	err := r.Db.Preload("Cliente").Preload("Arquivos").Preload("Procedure").Preload("ValorProposta").Find(&orcamento).Where("id = ? and status != ?", id, status).First(&orcamento).Error

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
		"items":    orcamento,
	})
}
