package models

import (
	"fiber/internal/configs"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/mobilemindtec/go-payments/asaas"
	"gorm.io/gorm"
)

type ProposedValue struct {
	gorm.Model
	Price    string `json:"price"`
	Amount   string `json:"amount"`
	Addition string `json:"addition"`
	Discount string `json:"discount"`
	X        string `json:"x"`
}

type Files struct {
	gorm.Model
	Url      string `json:"url"`
	Filename string `json:"filename"`
}

type Customer struct {
	gorm.Model
	Nome                string  `json:"nome"`
	DataNascimento      string  `json:"data_nascimento"`
	Cpf                 string  `json:"cpf" gorm:"unique;"`
	Rg                  string  `json:"rg"`
	Email               string  `json:"email" gorm:"unique" validate:"email" structs:"email"`
	Idade               int     `json:"idade"`
	Foto                string  `json:"foto"`
	EstadoCivil         string  `json:"estado_civil"`
	Sexo                string  `json:"sexo"`
	Contato             string  `json:"contato"`
	Contato2            string  `json:"contato2"`
	Cep                 string  `json:"cep"`
	Logradouro          string  `json:"logradouro"`
	Numero              string  `json:"numero"`
	Complemento         string  `json:"complemento"`
	Bairro              string  `json:"bairro"`
	Cidade              string  `json:"cidade"`
	Estado              string  `json:"estado"`
	Prontuario          string  `json:"prontuario"`
	Assasid             string  `json:"assas_id"`
	Situacao            string  `json:"situacao"`
	Indicacao           string  `json:"indicacao"`
	Profissao           string  `json:"profissao"`
	Observacao          string  `json:"observacao"`
	ConsultasCreditos   int     `json:"consultas_creditos"`
	ConsultasRealizadas int     `json:"consultas_realizadas"`
	ConsultasRestantes  int     `json:"consultas_restantes"`
	Midia               []Files `gorm:"many2many:customer_midias;"  json:"midias"`
	NomeResponsavel     string  `json:"nome_responsavel"`
	CpfResponsavel      string  `json:"cpf_responsavel"`
}

func (u *Customer) BeforeCreate(db *gorm.DB) (err error) {
	var last_id int

	if len(u.Prontuario) == 0 {
		err = db.Raw("SELECT MAX(id)+1 AS id FROM Customers").Scan(&last_id).Error
		if err != nil {
			return err
		}

		u.Prontuario = string(last_id)
	}

	token := configs.GetAsaasToken()
	pay := asaas.NewAsaas("BRL", token.AsaasToken, token.AsaasMode)

	resp, err := pay.CustomerCreate(&asaas.Customer{
		Name:                 u.Nome,
		CpfCnpj:              u.Cpf,
		Email:                u.Email,
		Phone:                u.Contato2,
		MobilePhone:          u.Contato,
		Address:              u.Logradouro,
		AddressNumber:        u.Numero,
		Province:             u.Bairro,
		PostalCode:           u.Cep,
		City:                 u.Cidade,
		NotificationDisabled: false,
		ExternalReference:    u.Prontuario,
	})

	if err != nil {
		fmt.Println(err)
	}

	u.Assasid = resp.Id

	return nil
}

type Groups struct {
	gorm.Model
	Nome string `json:"nome" validate:"required"`
}

type Procedure struct {
	gorm.Model
	Name     string `json:"nome" validate:"required"`
	Price    string `json:"preco" validate:"required"`
	Category string `json:"categoria" validate:"required"`
}

type Data struct {
	gorm.Model
	Dia int `json:"dia" validate:"required"`
	Mes int `json:"mes" validate:"required"`
	Ano int `json:"ano" validate:"required"`
}

type Tooth struct {
	gorm.Model
	Nome      string      `json:"nome"`
	Numero    string      `json:"numero"`
	Procedure []Procedure `gorm:"many2many:thooth_procedimentos;" json:"procedimentos"`
}

type Budget struct {
	gorm.Model
	Data               string          `json:"data" validate:"required"`
	Situacao           string          `json:"situacao"`
	Anotacoes          string          `json:"anotacoes"`
	FormaPagamento     string          `json:"forma_pagamento" validate:"required"`
	ClienteRefer       int             `json:"-"`
	Cliente            Customer        `gorm:"foreignKey:ClienteRefer;"  json:"cliente"`
	VendedorRefer      int             `json:"-"`
	Vendedor           User            `gorm:"foreignKey:VendedorRefer;"  json:"vendedor"`
	Arquivos           []Files         `gorm:"many2many:budget_arquivos;" json:"arquivos"`
	Tooth              []Tooth         `gorm:"many2many:budget_tooths;" json:"dentes_procedimento"`
	ValorProposta      []ProposedValue `gorm:"many2many:budget_propostas;" json:"valores_proposta"`
	Paymentid          string          `json:"paymentid"`
	ValorTotal         float64         `json:"valor_total" validate:"required"`
	Linkpagamento      string          `json:"link_pagamento"`
	InvoiceUrl         string          `json:"link_nota"`
	BankSlipUrl        string          `json:"link_boleto"`
	NetValue           float64         `json:"valor_liquido"`
	Quantidadeparcelas int             `json:"quantidade_parcelas"`
	ValorParcelas      int64           `json:"valor_parcelas"`
}

func (u *Budget) BeforeCreate(db *gorm.DB) (err error) {
	// var payment *asaas.Payment

	// if u.FormaPagamento == "BOLETO" {
	// 	var cliente Customer

	// 	err = db.Find(&cliente, u.ClienteRefer).Error

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	token := configs.GetAsaasToken()
	// 	pay := asaas.NewAsaas("BRL", token.AsaasToken, token.AsaasMode)

	// 	ChargeType := asaas.Detached
	// 	if u.Quantidadeparcelas > 1 {
	// 		ChargeType = asaas.Installment
	// 		payment.InstallmentCount = int64(u.Quantidadeparcelas)
	// 		payment.InstallmentValue = u.ValorParcelas
	// 		payment.TotalValue = u.ValorTotal
	// 	} else {
	// 		payment.Value = u.ValorTotal
	// 	}

	// 	payment = &asaas.Payment{
	// 		BillingType:       asaas.BillingType("BOLETO"),
	// 		DueDate:           u.Data,
	// 		Description:       "Denshow - Orçamento",
	// 		ExternalReference: u.Cliente.Prontuario,
	// 		PostalService:     false,
	// 		Name:              u.Cliente.Nome,
	// 		DueDateLimitDays:  5,
	// 		// ChargeType:        ChargeType,
	// 		Customer:            u.Cliente.Assasid,
	// 		NextDueDate:         u.Data,
	// 		SubscriptionCycle:   api.SubscriptionCycle("1"),
	// 		MaxInstallmentCount: int64(u.Quantidadeparcelas),
	// 	}

	// 	resp, err := pay.PaymentCreate(payment)

	// 	fmt.Println(resp, err)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	u.Paymentid = resp.Id
	// 	u.Situacao = "PENDING"
	// 	u.Data = resp.DateCreated
	// 	u.NetValue = resp.NetValue
	// 	u.Linkpagamento = resp.InvoiceUrl
	// 	u.BankSlipUrl = resp.BankSlipUrl
	// }

	u.Situacao = "PENDING"
	u.NetValue = u.ValorTotal

	return nil
}

type User struct {
	gorm.Model
	Username  string   `json:"username" validate:"required" gorm:"unique"`
	FirstName string   `json:"first_name" validate:"required"`
	LastName  string   `json:"last_name" validate:"required"`
	Email     string   `json:"email" gorm:"unique" validate:"email,omitempty,required" structs:"email,omitempty"`
	Password  string   `json:"password" validate:"required,min=6"`
	IsStaff   bool     `json:"is_staff" validate:"required"`
	IsActive  bool     `json:"is_active"`
	Groups    []Groups `gorm:"many2many:user_groups;" json:"grupos"`
}

var validate = validator.New()

func ValidateStruct(item interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(item)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
