package customer

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var customer models.Customer
	var foo models.Customer

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&customer).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Customer not found",
		})
		return err
	}

	err = r.Db.Model(&customer).UpdateColumns(models.Customer{
		Nome:                foo.Nome,
		Email:               foo.Email,
		DataNascimento:      foo.DataNascimento,
		Cpf:                 foo.Cpf,
		Rg:                  foo.Rg,
		Idade:               foo.Idade,
		Contato:             foo.Contato,
		Contato2:            foo.Contato2,
		EstadoCivil:         foo.EstadoCivil,
		Profissao:           foo.Profissao,
		Foto:                foo.Foto,
		Sexo:                foo.Sexo,
		Cep:                 foo.Cep,
		Logradouro:          foo.Logradouro,
		Numero:              foo.Numero,
		Complemento:         foo.Complemento,
		Bairro:              foo.Bairro,
		Cidade:              foo.Cidade,
		Estado:              foo.Estado,
		Prontuario:          foo.Prontuario,
		Situacao:            foo.Situacao,
		Indicao:             foo.Indicao,
		Observacao:          foo.Observacao,
		ConsultasCreditos:   foo.ConsultasCreditos,
		ConsultasRealizadas: foo.ConsultasRealizadas,
		ConsultasRestantes:  foo.ConsultasRestantes,
		Midia:               foo.Midia,
	}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Customer updated successfully",
		"item":    customer,
	})
	return nil
}
