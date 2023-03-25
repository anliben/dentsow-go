package utils

import (
	"fiber/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetCaixaEnd(app *fiber.Ctx) error {

	// models
	// var user models.User
	// var customer models.Customer
	var orcamento []models.Budget

	// params
	mes := app.Params("mes")
	ano := app.Params("ano")

	err := r.Db.Preload("Cliente").Preload("Vendedor").Preload("Data").Preload("Arquivos").Preload("Procedure").Preload("ValorProposta").Where("EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?", ano, mes).Find(&orcamento).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	valor_total := ""

	for i, item := range orcamento {
		valor_total += item.ValorProposta[i].Price
		fmt.Println(item.ValorProposta[i].Price)
	}

	return app.JSON(&fiber.Map{
		"count": len(orcamento),
		"items": orcamento,
	})
}
