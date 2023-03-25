package procedure

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

	routes := app.Group("/api/v1/procedure")
	routes.Post("/", r.Create)
	routes.Get("/", r.GetAll)
	routes.Get("/:id", r.GetById)
	routes.Put("/:id", r.Update)
	routes.Delete("/:id", r.Delete)
}
