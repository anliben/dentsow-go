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
	"fmt"
	"os"

	"github.com/eduardo-mior/mercadopago-sdk-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/mobilemindtec/go-payments/api"
	"github.com/mobilemindtec/go-payments/asaas"
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
	// db.AutoMigrate(&models.Groups{})
	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Files{})
	// db.AutoMigrate(&models.Customer{})
	// db.AutoMigrate(&models.Procedure{})
	// db.AutoMigrate(&models.ProposedValue{})
	// db.AutoMigrate(&models.Budget{})

	users.RegisterRoutes(app, db)
	customer.RegisterRoutes(app, db)
	groups.RegisterRoutes(app, db)
	proposed.RegisterRoutes(app, db)
	procedure.RegisterRoutes(app, db)
	files.RegisterRoutes(app, db)
	utils.RegisterRoutes(app, db)
	budget.RegisterRoutes(app, db)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Post("/webhook/mercadopago", func(c *fiber.Ctx) error {
		var orcamento models.Budget

		// pegar corpo da requisição
		var webhookResponse mercadopago.WebhookResponse

		err := c.BodyParser(&webhookResponse)

		response, mercadopagoErr, err := mercadopago.ConsultPayment(webhookResponse.Data.ID, "TEST-3692262666358677-033011-1bf16959d504fa3072556d236bc3134f-425659019")

		if err != nil {
			// Erro inesperado
		} else if mercadopagoErr != nil {
			// Erro retornado do MercadoPago
		} else {
			// Sucesso!
		}

		err = db.First(&orcamento).UpdateColumns(models.Budget{
			Situacao: response.Status,
		}).Where("Paymentid = ?", webhookResponse.Data.ID).Error

		if err != nil {
			return c.JSON(&fiber.Map{
				"status": "error",
			})
		}
		return nil
	})

	mercado()

	app.Listen(getPort())
}

func mercado() {
	//
}

func runner() {
	pay := asaas.NewAsaas("", "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAwNTIxMTY6OiRhYWNoXzJiN2M1YzI0LTNmYjktNDE4Ni04NmM3LTQzNzUxYzhjNGFhYw==", api.AsaasModeTest)

	resp, _ := pay.PaymentCreate(&asaas.Payment{
		BillingType:       "UNDEFINED",
		Value:             250,
		Description:       "clareamento de dentes",
		Name:              "Joao victor paulino silva",
		DueDateLimitDays:  1,
		DueDate:           "2023-09-01",
		ChargeType:        "DETACHED",
		Customer:          "5208495",
		ExternalReference: "123",
		NextDueDate:       "2023-09-01",
		SubscriptionCycle: api.SubscriptionCycle(1),
	})

	fmt.Println(resp)
}
