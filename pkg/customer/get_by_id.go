package customer

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetById(app *fiber.Ctx) error {
	var customer models.Customer

	err := r.Db.Find(&customer).Where("id = ?", app.Params("id")).First(&customer).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    1,
		"next":     "null",
		"previous": "null",
		"items":    customer,
	})
}
