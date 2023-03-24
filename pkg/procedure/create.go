package procedure

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var procedure models.Procedure

	err := app.BodyParser(&procedure)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Create(&procedure).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Procedure created successfully",
		"item":    procedure,
	})
	return nil
}
