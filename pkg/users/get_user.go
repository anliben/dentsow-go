package users

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r *handler) UserGetAll(app *fiber.Ctx) error {
	var users []models.User

	err := r.Db.Preload("Groups").Preload("Permissions").Find(&users).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Users not found",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(users),
		"next":     "null",
		"previous": "null",
		"items":    users,
	})
}
