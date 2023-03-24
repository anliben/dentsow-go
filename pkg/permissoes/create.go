package permissoes

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var permissions models.Permissions

	err := app.BodyParser(&permissions)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Create(&permissions).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Permission created successfully",
		"item":    permissions,
	})
	return nil
}
