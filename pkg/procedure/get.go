package procedure

import (
	"fiber/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (r *handler) GetAll(app *fiber.Ctx) error {
	var procedure []models.Procedure
	r.Db.Find(&procedure)

	return app.JSON(&fiber.Map{
		"count":    len(procedure),
		"next":     "null",
		"previous": "null",
		"items":    procedure,
	})
}
