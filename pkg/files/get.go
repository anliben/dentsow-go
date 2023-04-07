package files

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var files []models.Files
	
	tx := r.Db.Session(&gorm.Session{PrepareStmt: true})

	err := tx.
		Find(&files).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(files),
		"next":     "null",
		"previous": "null",
		"items":    files,
	})
}
