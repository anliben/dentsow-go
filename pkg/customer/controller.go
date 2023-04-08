package customer

import (
	"fiber/pkg/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	routes.Use(users.AuthMiddleware, cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))
	routes.Get("/", r.GetAll)
	routes.Post("/", r.Create)
	routes.Get("/:id", r.GetById)
	routes.Put("/:id", r.Update)
	routes.Delete("/:id", r.Delete)
}
