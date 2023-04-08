package users

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

var (
	store   *session.Store
	REFRESH string = "refresh"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	r := &handler{
		Db: db,
	}

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 160,
	})

	routes := app.Group("/api/v1/users")
	// routes.Use(AuthMiddleware)
	routes.Get("/", r.UserGetAll)
	routes.Get("/:id", r.GetById)
	routes.Post("/", r.UserCreateOne)
	routes.Delete("/:id", r.Delete)
	routes.Put("/:id", r.Update)
	routes.Post("/jwt/create", r.Sign)
	routes.Get("/jwt/verify", r.HealthCheck)
	routes.Get("/jwt/logout", r.Logout)
}

func AuthMiddleware(c *fiber.Ctx) error {

	sess, err := store.Get(c)

	if strings.Split(c.Path(), "/")[1] == "jwt/create" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	if sess.Get(REFRESH) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}
