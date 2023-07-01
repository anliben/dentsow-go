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
	"fmt"
	"time"

	//"fiber/pkg/common/models"

	_ "fiber/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func Setup() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4553,https://dentshow-ui.vercel.app",
		AllowHeaders:     "",
		AllowCredentials: true,
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

	app.Get("/swagger/*", fiberSwagger.WrapHandler) // default

	// go Hello()

	utils.StartServerWithGracefulShutdown(app)

}

func Hello() {
	for index := 0; index < 10000; index++ {
		fmt.Println("Estou ainda rodando em segundo plano")
		time.Sleep(time.Millisecond * 250)
	}
}
