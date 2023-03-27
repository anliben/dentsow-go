package customer

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var customer []models.Customer
	err := r.Db.Preload("Midia").Find(&customer).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(customer),
		"next":     "null",
		"previous": "null",
		"items":    customer,
	})
}