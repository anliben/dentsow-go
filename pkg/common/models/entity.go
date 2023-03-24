package models

import (
	"time"

	"gorm.io/gorm"
)

type ProposedValue struct {
	gorm.Model
	Price    string `json:"price"`
	Amount   int    `json:"amount"`
	Addition int    `json:"addition"`
	Discount int    `json:"discount"`
}

type Files struct {
	gorm.Model
	Url string `json:"url"`
}

type Customer struct {
	Id                  int     `json:"id"`
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
	Situacao            string  `json:"situacao"`
	Indicao             string  `json:"indicao"`
	Profissao           string  `json:"profissao"`
	Observacao          string  `json:"observacao"`
	ConsultasCreditos   int     `json:"consultas_creditos"`
	ConsultasRealizadas int     `json:"consultas_realizadas"`
	ConsultasRestantes  int     `json:"consultas_restantes"`
	Midia               []Files `gorm:"many2many:customer_midias;"`
}

type Groups struct {
	gorm.Model
	Name string `json:"name"`
}

type Permissions struct {
	gorm.Model
	Name     string `json:"name"`
	CodeName string `json:"codename"`
}

type Procedure struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Category string `json:"category"`
}

type Budget struct {
	Id             int           `json:"id"`
	ValorProposta  ProposedValue `json:"valor_proposta"`
	Data           string        `json:"data"`
	Situacao       string        `json:"situacao"`
	Anotacoes      string        `json:"anotacoes"`
	FormaPagamento string        `json:"forma_pagamento"`
	Created        time.Time     `json:"created"`
	Updated        time.Time     `json:"updated"`
	Cliente        int           `json:"cliente"`
	Arquivos       []Files       `gorm:"many2many:budget_arquivos;"`
}

type User struct {
	gorm.Model
	Username    string        `json:"username"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	IsStaff     bool          `json:"is_staff"`
	IsActive    bool          `json:"is_active"`
	Permissions []Permissions `gorm:"many2many:user_permissions;"`
	Groups      []Groups      `gorm:"many2many:user_groups;"`
}
