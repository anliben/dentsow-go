package users

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

	routes := app.Group("/api/v1/users")
	routes.Get("/", r.UserGetAll)
	routes.Get("/:id", r.GetById)
	routes.Post("/", r.UserCreateOne)
	routes.Delete("/:id", r.Delete)
	routes.Put("/:id", r.Update)
	routes.Post("/jwt/create", r.Sign)
}
