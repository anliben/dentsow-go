package utils

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetCountIdTable(app *fiber.Ctx) error {

	// var user models.User

	table := app.Params("table")

	if table == "" {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Table is empty",
		})
		return nil
	}

	err := r.Db.Raw(fmt.Sprintf("SELECT MAX(id+1) FROM %s", table)).Scan(&table).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"next_id": table,
	})
}
