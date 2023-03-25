package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetById(app *fiber.Ctx) error {
	var orcamento models.Budget

	err := r.Db.Preload("Cliente").Preload("Arquivos").Preload("Procedure").Preload("ValorProposta").Find(&orcamento).Where("id = ?", app.Params("id")).First(&orcamento).Error

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
