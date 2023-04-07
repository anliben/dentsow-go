package proposed

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r handler) Create(app *fiber.Ctx) error {
	var proposed models.ProposedValue

	err := app.BodyParser(&proposed)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	errors := models.ValidateStruct(proposed)
	if errors != nil {
		return app.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = r.Db.Create(&proposed).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "User created successfully",
		"item":    proposed,
	})
	return nil
}
