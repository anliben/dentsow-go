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

	//"fiber/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4553,https://dentshow-web.up.railway.app",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 2
	}))

	db, _ := database.OpenConnection()

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Budget{})
	// db.AutoMigrate(&models.Customer{})
	// db.AutoMigrate(&models.Data{})
	// db.AutoMigrate(&models.Files{})
	// db.AutoMigrate(&models.Groups{})
	// db.AutoMigrate(&models.Procedure{})
	// db.AutoMigrate(&models.ProposedValue{})

	users.RegisterRoutes(app, db)
	customer.RegisterRoutes(app, db)
	groups.RegisterRoutes(app, db)
	proposed.RegisterRoutes(app, db)
	procedure.RegisterRoutes(app, db)
	files.RegisterRoutes(app, db)
	utils.RegisterRoutes(app, db)
	budget.RegisterRoutes(app, db)

	utils.StartServerWithGracefulShutdown(app)

}
