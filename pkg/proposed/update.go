package proposed

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var proposta models.ProposedValue
	var foo models.ProposedValue

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&proposta).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Proposta not found",
		})
		return err
	}

	err = r.Db.Model(&proposta).UpdateColumns(models.ProposedValue{
		Price:    foo.Price,
		Amount:   foo.Amount,
		Addition: foo.Addition,
		Discount: foo.Discount,
	}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Proposta updated successfully",
		"item":    proposta,
	})
	return nil
}
