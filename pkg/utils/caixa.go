package utils

import (
	"fiber/pkg/common/models"
	"net/http"

	"strconv"

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
	X                 string
	Juros             float64
	Liquido           float64
	TipoPagamento     string
	Descricao         string
	Vencimento        string
	UrlRecebimento    string
	DataDaConfirmacao string
	Discount          string
	Multas            string
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

	err := r.Db.
		Preload("Cliente").
		Preload("Vendedor").
		Preload("Arquivos").
		Preload("Procedure").
		Preload("ValorProposta").
		Where("EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?", ano, mes).
		Find(&orcamento).Error

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
		if item.FormaPagamento != "BOLETO" {

			var discount float64

			for _, item := range item.ValorProposta {
				marks, _ := strconv.Atoi(item.Discount)

				if marks != 0 {
					discount += float64(marks)
				}
			}

			caixa := Caixa{
				Data:              item.Data,
				Prontuario:        item.Cliente.Prontuario,
				X:                 item.ValorProposta[index].X,
				Status:            orcamento[index].Situacao,
				TipoPagamento:     string(item.FormaPagamento),
				Descricao:         item.Anotacoes,
				Vencimento:        item.Data,
				DataDaConfirmacao: item.Data,
				ClienteAssas:      item.Cliente.Assasid,
				Total:             item.ValorTotal,
				Juros:             0,
				Liquido:           item.NetValue,
				Discount:          strconv.FormatFloat(discount, 'E', -1, 64),
				Multas:            "sem multas",
			}

			listCaixa = append(listCaixa, caixa)
		} else {
			p, _ := pay.PaymentGet(item.Paymentid)
			// for index, item := range item.Procedure {
			// 	fmt.Println(index, item)
			// }

			caixa := Caixa{
				Data:              item.Data,
				Prontuario:        item.Cliente.Prontuario,
				X:                 strconv.FormatInt(p.InstallmentCount, 10),
				Status:            orcamento[index].Situacao,
				TipoPagamento:     string(p.BillingType),
				Descricao:         p.Description,
				Vencimento:        p.OriginalDueDate,
				DataDaConfirmacao: p.ConfirmedDate,
				ClienteAssas:      p.Customer,
				Total:             item.ValorTotal,
				Juros:             p.Interest.Value,
				Liquido:           p.NetValue,
				Discount:          strconv.FormatFloat(p.Discount.Value, 'E', -1, 64),
				Multas:            strconv.FormatFloat(p.Fine.Value, 'E', -1, 64),
			}

			listCaixa = append(listCaixa, caixa)
		}
	}

	return app.JSON(&fiber.Map{
		"count":           len(orcamento),
		"items":           listCaixa,
		"paymentNotFound": listPayNotFound,
	})
}
