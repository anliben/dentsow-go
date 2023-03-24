package users

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r handler) GetById(app *fiber.Ctx) error {
	var user models.User

	id := app.Params("id")

	if id == "" {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID is empty",
		})
		return nil
	}

	err := r.Db.Preload("Groups").Preload("Permissions").Where("id = ?", id).First(&user).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "User not found",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"item": user,
	})
}
