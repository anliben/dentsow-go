package customer

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @title Fiber Swagger Example API
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Customer
// @Router / [get]
func (r handler) GetAll(app *fiber.Ctx) error {
	var customer []models.Customer

	tx := r.Db.Session(&gorm.Session{PrepareStmt: true})

	err := tx.
		Preload("Midia").
		Find(&customer).Error

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
