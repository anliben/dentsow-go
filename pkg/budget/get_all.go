package budget

import (
	"fiber/pkg/common/models"
	"fiber/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var budget []models.Budget

	data := app.Query("data")
	situacao := app.Query("situacao")
	anotacoes := app.Query("anotacoes")
	forma_pagamento := app.Query("forma_pagamento")
	valor_total := app.Query("valor_total")
	id := app.Query("id")


	db := r.Db.
		Preload("Cliente").
		Preload("Cliente.Midia").
		Preload("Vendedor").
		Preload("Vendedor.Groups").
		Preload("Arquivos").
		Preload("Tooth").
		Preload("Tooth.Procedure").
		Preload("ValorProposta")

		if id != "" {
			err := db.First(&budget, id).Error
			
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
				"items":    budget,})
		}
	

	db.Where(utils.Builder("data LIKE ?", "%"+data+"%"))
	db.Where(utils.Builder("situacao LIKE ?", "%"+situacao+"%"))
	db.Where(utils.Builder("anotacoes LIKE ?", "%"+anotacoes+"%"))
	db.Where(utils.Builder("forma_pagamento LIKE ?", "%"+forma_pagamento+"%"))
	db.Where(utils.Builder("valor_total = ?", valor_total))

	err := db.Find(&budget).Error

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
