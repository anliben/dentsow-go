package users

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	errors := models.ValidateStruct(user)
	if errors != nil {
		return app.Status(fiber.StatusBadRequest).JSON(errors)

	}

	password, bcErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcErr != nil {
		err = app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Erro ao criptografar senha!",
			"error":  err.Error(),
		})
		return err
	}

	user.Password = string(password)

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
