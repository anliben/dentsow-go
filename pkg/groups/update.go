package groups

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var grupo models.Groups
	var foo models.Groups

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&grupo).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Grupos not found",
		})
		return err
	}

	err = r.Db.Model(&grupo).UpdateColumns(models.Groups{
		Nome: foo.Nome,
	}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Grupos updated successfully",
		"item":    grupo,
	})
	return nil
}
