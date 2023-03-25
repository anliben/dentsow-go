package main

import (
	"fiber/internal/database"
	"fiber/pkg/budget"
	"fiber/pkg/common/models"
	"fiber/pkg/customer"
	"fiber/pkg/files"
	"fiber/pkg/groups"
	"fiber/pkg/procedure"
	"fiber/pkg/proposed"
	"fiber/pkg/users"
	"fiber/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	app.Use(recover.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	db, _ := database.OpenConnection()
	db.AutoMigrate(&models.Groups{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Files{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Procedure{})
	db.AutoMigrate(&models.ProposedValue{})
	db.AutoMigrate(&models.Budget{})

	users.RegisterRoutes(app, db)
	customer.RegisterRoutes(app, db)
	groups.RegisterRoutes(app, db)
	proposed.RegisterRoutes(app, db)
	procedure.RegisterRoutes(app, db)
	files.RegisterRoutes(app, db)
	utils.RegisterRoutes(app, db)
	budget.RegisterRoutes(app, db)

	app.Listen(getPort())
}
