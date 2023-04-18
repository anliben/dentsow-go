package utils

import (
	"fiber/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mobilemindtec/go-payments/api"
	"github.com/mobilemindtec/go-payments/asaas"
)

type Caixa struct {
	Data              string
	Prontuario        string
	ClienteAssas      string
	Status            string
	Total             float64
	X                 int64
	Juros             float64
	Liquido           float64
	TipoPagamento     string
	Descricao         string
	Vencimento        string
	UrlRecebimento    string
	DataDaConfirmacao string
	Discount          float64
	Multas            float64
}

type PaymentNotFound struct {
	IDPaymentNotFound string
}

type CaixaFechamento struct {
	TotalDebitoLiquido float64
	TotalDebitoBruto   float64

	TotalCartaoLiquido float64
	TotalCartaoBruto   float64
	TotalCartaoTaxa    float64

	TotalDinheiro float64
	TotalLiquido  float64

	QuantidadeCartao   int64
	QuantidadePix      int64
	QuantidadeDinheiro int64

	TotalReceber int64
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

	listCaixa := []Caixa{}
	listPayNotFound := []PaymentNotFound{}

	pay := asaas.NewAsaas("BRL", "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAyOTAxMTE6OiRhYWNoX2EzZGNmMDY1LWM0MWYtNDg4OC05ZjNlLTRmOGVlNTczMjQyMw==", api.AsaasModeProd)

	for index, item := range orcamento {
		p, _ := pay.PaymentGet(item.Paymentid)

		for index, item := range item.Procedure {
			fmt.Println(index, item)
		}

		caixa := Caixa{
			Data:              item.Data,
			Prontuario:        item.Cliente.Prontuario,
			X:                 p.InstallmentCount,
			Status:            orcamento[index].Situacao,
			TipoPagamento:     string(p.BillingType),
			Descricao:         p.Description,
			Vencimento:        p.OriginalDueDate,
			DataDaConfirmacao: p.ConfirmedDate,
			ClienteAssas:      p.Customer,
			Total:             item.ValorTotal,
			Juros:             p.Interest.Value,
			Liquido:           p.NetValue,
			Discount:          p.Discount.Value,
			Multas:            p.Fine.Value,
		}

		listCaixa = append(listCaixa, caixa)
	}

	return app.JSON(&fiber.Map{
		"count":           len(orcamento),
		"items":           listCaixa,
		"paymentNotFound": listPayNotFound,
	})
}
