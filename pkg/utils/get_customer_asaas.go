package utils

import (
	"fiber/internal/configs"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mobilemindtec/go-payments/asaas"
)

func (r handler) GetCustomerList(app *fiber.Ctx) error {

	token := configs.GetAsaasToken()
	handle := asaas.NewAsaas("BRL", token.AsaasToken, token.AsaasMode)

	customers_response, _ := handle.CustomerFind(asaas.NewDefaultFilter().ToMap())

	customers := customers_response.CustomerResults.Data

	for customer := range customers {
		fmt.Println(customer)
	}

	return nil
}
