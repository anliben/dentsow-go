package procedure

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var procedimento models.Procedure
	var foo models.Procedure

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&procedimento).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Procedimento not found",
		})
		return err
	}

	err = r.Db.Model(&procedimento).UpdateColumns(models.Procedure{
		Name:     foo.Name,
		Price:    foo.Price,
		Category: foo.Category,
	}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Procedimento updated successfully",
		"item":    procedimento,
	})
	return nil
}
