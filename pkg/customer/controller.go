package customer

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

	routes := app.Group("/api/v1/clientes")
	routes.Get("/", r.GetAll)
	routes.Get("/:id", r.GetById)
	routes.Post("/", r.Create)
	routes.Put("/:id", r.Update)
	routes.Delete("/:id", r.Delete)
}
