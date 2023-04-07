package customer

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var customer models.Customer

	err := app.BodyParser(&customer)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	errors := models.ValidateStruct(customer)
	if errors != nil {
		return app.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = r.Db.Create(&customer).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "User created successfully",
		"item":    customer,
	})
	return nil
}
