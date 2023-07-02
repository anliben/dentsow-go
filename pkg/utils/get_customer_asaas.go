package utils

import (
	"fiber/internal/configs"
	"fiber/pkg/common/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mobilemindtec/go-payments/asaas"
)

func (r handler) GetCustomerList(app *fiber.Ctx) error {
	var customer_response models.Customer

	token := configs.GetAsaasToken()
	handle := asaas.NewAsaas("BRL", token.AsaasToken, token.AsaasMode)

	customers_response, _ := handle.CustomerFind(asaas.NewDefaultFilter().ToMap())

	customers := customers_response.CustomerResults.Data

	for _, c := range customers {
		db := r.Db.Where(Builder("Assasid = ?", c.Id))

		err := db.Find(&customer_response).Error
		if err != nil {

			cus := &models.Customer{
				Nome:       c.Name,
				Cpf:        c.CpfCnpj,
				Email:      c.Email,
				Contato:    c.MobilePhone,
				Contato2:   c.Phone,
				Logradouro: c.Address,
				Numero:     c.AddressNumber,
				Bairro:     c.Province,
				Cep:        c.PostalCode,
				Assasid:    c.Id,
			}

			// errors := models.ValidateStruct(cus)
			// if errors != nil {
			// 	return app.Status(fiber.StatusBadRequest).JSON(errors)
			// }
			err := r.Db.Create(cus).Error

			if err != nil {
				fmt.Println(err)
			}
		}

	}

	return nil
}
