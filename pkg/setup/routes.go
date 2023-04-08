package setup

import (
	"encoding/json"
	"fiber/internal/database"
	"fiber/pkg/budget"
	"fiber/pkg/customer"
	"fiber/pkg/files"
	"fiber/pkg/groups"
	"fiber/pkg/procedure"
	"fiber/pkg/proposed"
	"fiber/pkg/users"
	"fiber/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup() error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(users.AuthMiddleware, cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 2
	}))
	db, _ := database.OpenConnection()

	users.RegisterRoutes(app, db)
	customer.RegisterRoutes(app, db)
	groups.RegisterRoutes(app, db)
	proposed.RegisterRoutes(app, db)
	procedure.RegisterRoutes(app, db)
	files.RegisterRoutes(app, db)
	utils.RegisterRoutes(app, db)
	budget.RegisterRoutes(app, db)

	err := app.Listen(getPort())
	return err
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}
