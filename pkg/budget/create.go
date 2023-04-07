package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var orcamento models.Budget

	err := app.BodyParser(&orcamento)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	errors := models.ValidateStruct(orcamento)
	if errors != nil {
		return app.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = r.Db.Create(&orcamento).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Orcamento created successfully",
		"item":    orcamento,
	})
	return nil
}
