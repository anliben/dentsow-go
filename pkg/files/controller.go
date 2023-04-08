package files

import (
	"fiber/pkg/users"

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

	routes := app.Group("/api/v1/files")
	routes.Use(users.AuthMiddleware)
	routes.Post("/", r.Upload)
	routes.Get("/", r.GetAll)
}
