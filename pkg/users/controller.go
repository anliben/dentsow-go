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
		CookieSecure:   false,
		Expiration: time.Hour * 160,
	})

	routes := app.Group("/api/v1/users")
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

	if strings.Join(strings.Split(c.Path(), "/"), "/") == "/api/v1/users/jwt/create" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Atenção! Esta é uma zona restrita apenas para pessoas autorizadas. Se você não tem a senha secreta, é melhor ir pegar um café e tentar novamente mais tarde.",
		})
	}

	if sess.Get(REFRESH) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Atenção! Esta é uma zona restrita apenas para pessoas autorizadas. Se você não tem a senha secreta, é melhor ir pegar um café e tentar novamente mais tarde.",
		})
	}

	return c.Next()
}
