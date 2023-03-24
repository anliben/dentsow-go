package permissoes

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) PermissionGetAll(app *fiber.Ctx) error {
	var permisson []models.Permissions
	err := r.Db.Find(&permisson).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(permisson),
		"next":     "null",
		"previous": "null",
		"items":    permisson,
	})
}
