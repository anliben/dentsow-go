package tooth

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var tooths []models.Files

	tx := r.Db.Session(&gorm.Session{PrepareStmt: true})

	err := tx.
		Find(&tooths).Error

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
			"detail":  err.Error(),
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(tooths),
		"next":     "null",
		"previous": "null",
		"items":    tooths,
	})
}
