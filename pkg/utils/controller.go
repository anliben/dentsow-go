package utils

import (
	"fiber/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{
		Db: db,
	}

	db.AutoMigrate(&models.Groups{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Files{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Procedure{})
	db.AutoMigrate(&models.ProposedValue{})
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.Data{})

	routes := app.Group("/api/v1/utils")
	routes.Get("/:table", r.GetCountIdTable)
	routes.Get("/:mes/:ano", r.GetCaixaEnd)
	routes.Get("/migrate", r.Migrate)
}
