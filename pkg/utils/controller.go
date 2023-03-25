package utils

import (
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

	routes := app.Group("/api/v1/utils")
	routes.Get("/:table", r.GetCountIdTable)
	routes.Get("/:mes/:ano", r.GetCaixaEnd)
}
