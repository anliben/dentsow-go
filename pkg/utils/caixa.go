package utils

import (
	"fiber/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mobilemindtec/go-payments/api"
	"github.com/mobilemindtec/go-payments/asaas"
)

type Cartao struct {
	Maquina string
	X       string
	Juros   int
	Total   string
	Liquido string
	Status  string
}

type Pix struct {
	Total  float64
	Pago   float64
	Status string
}

type Caixa struct {
	Data       string
	Prontuario string
	Credito    Cartao
	Debito     Cartao
	Pix        Pix
}

func (r handler) GetCaixaEnd(app *fiber.Ctx) error {
	// models
	// var user models.User
	// var customer models.Customer
	var orcamento []models.Budget
	// params
	mes := app.Params("mes")
	ano := app.Params("ano")

	err := r.Db.Preload("Cliente").Preload("Vendedor").Preload("Arquivos").Preload("Procedure").Preload("ValorProposta").Where("EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?", ano, mes).Find(&orcamento).Error

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	caixa := Caixa{}
	listCaixa := []Caixa{}

	pay := asaas.NewAsaas("", "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAwNTIxMTY6OiRhYWNoXzJiN2M1YzI0LTNmYjktNDE4Ni04NmM3LTQzNzUxYzhjNGFhYw==", api.AsaasModeTest)

	for index, item := range orcamento {
		p, err := pay.PaymentGet(item.Paymentid)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		fmt.Println(p)

		if p.BillingType == "BOLETO" || p.BillingType == "PIX" || p.BillingType == "UNDEFINED" {
			caixa = Caixa{
				Data:       p.DueDate,
				Prontuario: "item.Cliente[0].Prontuario",
				Pix: Pix{
					Total:  p.OriginalValue,
					Pago:   p.TotalBalance,
					Status: p.StatusText,
				},
			}
		}

		if p.BillingType == "CREDIT_CARD" {
			caixa = Caixa{
				Data:       p.DueDate,
				Prontuario: "item.Cliente[index].Prontuario",
				Credito: Cartao{
					X:      item.ValorProposta[index].X,
					Juros:  item.ValorProposta[index].Addition,
					Total:  p.OriginalDueDate,
					Status: p.StatusText,
				},
			}
		}

		listCaixa = append(listCaixa, caixa)
	}

	fmt.Println(listCaixa)

	return app.JSON(&fiber.Map{
		"count": len(orcamento),
		"items": listCaixa,
	})
}
