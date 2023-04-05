package models

import (
	"fiber/internal/database"
	"fmt"

	"github.com/google/uuid"
	"github.com/mobilemindtec/go-payments/api"
	"github.com/mobilemindtec/go-payments/asaas"
	"gorm.io/gorm"
)

type ProposedValue struct {
	gorm.Model
	Price    string `json:"price"`
	Amount   int    `json:"amount"`
	Addition int    `json:"addition"`
	Discount int    `json:"discount"`
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
	Cpf                 string  `json:"cpf"`
	Rg                  string  `json:"rg"`
	Email               string  `json:"email"`
	Idade               int     `json:"idade"`
	Foto                string  `json:"foto"`
	EstadoCivil         string  `json:"estado_civil"`
	Sexo                string  `json:"sexo"`
	Celular             string  `json:"celular"`
	Telefone            string  `json:"telefone"`
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
	Indicao             string  `json:"indicao"`
	Profissao           string  `json:"profissao"`
	Observacao          string  `json:"observacao"`
	ConsultasCreditos   int     `json:"consultas_creditos"`
	ConsultasRealizadas int     `json:"consultas_realizadas"`
	ConsultasRestantes  int     `json:"consultas_restantes"`
	Midia               []Files `gorm:"many2many:customer_midias;"  json:"midias"`
}

func (u *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	if u.Prontuario == "" {
		u.Prontuario = uuid.String()
	}

	pay := asaas.NewAsaas("BRL", "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAyOTAxMTE6OiRhYWNoX2EzZGNmMDY1LWM0MWYtNDg4OC05ZjNlLTRmOGVlNTczMjQyMw==", api.AsaasModeProd)

	resp, err := pay.CustomerCreate(&asaas.Customer{
		Name:                 u.Nome,
		CpfCnpj:              u.Cpf,
		Email:                u.Email,
		Phone:                u.Celular,
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
	Nome string `json:"nome"`
}

type Procedure struct {
	gorm.Model
	Name     string `json:"nome"`
	Price    string `json:"preco"`
	Category string `json:"categoria"`
}

type Data struct {
	gorm.Model
	Dia int `json:"dia"`
	Mes int `json:"mes"`
	Ano int `json:"ano"`
}

type Budget struct {
	gorm.Model
	DataRefer      int             `json:"data_refer"`
	Data           string          `json:"data"`
	Situacao       string          `json:"situacao"`
	Anotacoes      string          `json:"anotacoes"`
	FormaPagamento string          `json:"forma_pagamento"`
	VendedorRefer  int             `json:"vendedor_referer"`
	ClienteRefer   int             `json:"cliente_refer"`
	Cliente        Customer        `gorm:"foreignKey:VendedorRefer;"  json:"cliente"`
	Vendedor       User            `gorm:"foreignKey:VendedorRefer;"  json:"vendedor"`
	Arquivos       []Files         `gorm:"many2many:budget_arquivos;" json:"arquivos"`
	Procedure      []Procedure     `gorm:"many2many:budget_orcamentos;" json:"procedimentos"`
	ValorProposta  []ProposedValue `gorm:"many2many:budget_propostas;" json:"valores_proposta"`
	Paymentid      string          `json:"paymentid"`
	ValorTotal     float64         `json:"valor_total"`
	Linkpagamento  string          `json:"link_pagamento"`
	InvoiceUrl     string          `json:"link_nota"`
	BankSlipUrl    string          `json:"link_boleto"`
	NetValue       float64         `json:"valor_liquido"`
}

func (u *Budget) BeforeCreate(tx *gorm.DB) (err error) {
	var cliente Customer
	db, _ := database.OpenConnection()

	err = db.Find(&cliente, u.ClienteRefer).Error

	if err != nil {
		fmt.Println(err)
	}

	pay := asaas.NewAsaas("", "$aact_YTU5YTE0M2M2N2I4MTliNzk0YTI5N2U5MzdjNWZmNDQ6OjAwMDAwMDAwMDAwMDAyOTAxMTE6OiRhYWNoX2EzZGNmMDY1LWM0MWYtNDg4OC05ZjNlLTRmOGVlNTczMjQyMw==", api.AsaasModeProd)

	resp, err := pay.PaymentCreate(&asaas.Payment{
		BillingType:       asaas.BillingType(u.FormaPagamento),
		Value:             u.ValorTotal,
		Description:       "Denshow - Orçamento",
		Name:              cliente.Nome,
		DueDateLimitDays:  5,
		DueDate:           u.Data,
		ChargeType:        "DETACHED",
		Customer:          cliente.Assasid,
		ExternalReference: cliente.Prontuario,
		NextDueDate:       u.Data,
		SubscriptionCycle: api.SubscriptionCycle(1),
	})

	if err != nil {
		return err
	}

	u.Paymentid = resp.Id
	u.Situacao = "PENDING"
	u.Data = resp.DateCreated
	u.NetValue = resp.NetValue
	u.Linkpagamento = resp.InvoiceUrl
	u.BankSlipUrl = resp.BankSlipUrl

	return err
}

type User struct {
	gorm.Model
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	IsStaff   bool     `json:"is_staff"`
	IsActive  bool     `json:"is_active"`
	Groups    []Groups `gorm:"many2many:user_groups;" json:"grupos"`
}
