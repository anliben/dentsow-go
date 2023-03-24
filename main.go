package main

import (
	"fiber/internal/database"
	"fiber/pkg/customer"
	"fiber/pkg/users"
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
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Budget{})
	// db.AutoMigrate(&models.Customer{})
	// db.AutoMigrate(&models.Groups{})
	// db.AutoMigrate(&models.Permissions{})
	// db.AutoMigrate(&models.Procedure{})
	// db.AutoMigrate(&models.ProposedValue{})

	users.RegisterRoutes(app, db)
	customer.RegisterRoutes(app, db)

	app.Listen(getPort())
}
