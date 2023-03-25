package users

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var user models.User
	var foo models.User

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&user).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "User not found",
		})
		return err
	}

	err = r.Db.Model(&user).UpdateColumns(models.User{
		FirstName: foo.FirstName,
		LastName:  foo.LastName,
		Email:     foo.Email,
		Password:  foo.Password,
		Username:  foo.Username,
		IsStaff:   foo.IsStaff,
		IsActive:  foo.IsActive,
		Groups:    foo.Groups,
	}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "User updated successfully",
		"item":    user,
	})
	return nil
}
