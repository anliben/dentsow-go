package users

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r handler) UserCreateOne(app *fiber.Ctx) error {
	var user models.User

	err := app.BodyParser(&user)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Create(&user).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "User created successfully",
		"item":    user,
	})
	return nil
}
