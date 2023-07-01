package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var orcamento models.Budget
	var foo models.Budget
	var customer models.Customer
	var vendedor models.User

	err := app.BodyParser(&foo)
	id := app.Params("id")

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "Orcamento invalido",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Where("id = ?", id).First(&customer).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "Cliente nao encontrado!",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.Model(&customer).UpdateColumns(models.Customer{
		Nome:                foo.Cliente.Nome,
		DataNascimento:      foo.Cliente.DataNascimento,
		Cpf:                 foo.Cliente.Cpf,
		Rg:                  foo.Cliente.Rg,
		Email:               foo.Cliente.Email,
		Idade:               foo.Cliente.Idade,
		Foto:                foo.Cliente.Foto,
		EstadoCivil:         foo.Cliente.EstadoCivil,
		Sexo:                foo.Cliente.Sexo,
		Contato:             foo.Cliente.Contato,
		Contato2:            foo.Cliente.Contato2,
		Cep:                 foo.Cliente.Cep,
		Logradouro:          foo.Cliente.Logradouro,
		Numero:              foo.Cliente.Numero,
		Complemento:         foo.Cliente.Complemento,
		Bairro:              foo.Cliente.Bairro,
		Cidade:              foo.Cliente.Cidade,
		Estado:              foo.Cliente.Estado,
		Prontuario:          foo.Cliente.Prontuario,
		Assasid:             foo.Cliente.Assasid,
		Situacao:            foo.Cliente.Situacao,
		Indicacao:             foo.Cliente.Indicacao,
		Profissao:           foo.Cliente.Profissao,
		Observacao:          foo.Cliente.Observacao,
		ConsultasCreditos:   foo.Cliente.ConsultasCreditos,
		ConsultasRealizadas: foo.Cliente.ConsultasRealizadas,
		ConsultasRestantes:  foo.Cliente.ConsultasRestantes,
		Midia:               foo.Cliente.Midia,
	}).Where("id = ?", foo.Cliente.ID).Error

	err = r.Db.Where("id = ?", id).First(&vendedor).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "Vendedor nao encontrado!",
			"error":  err.Error(),
		})
		return err
	}

	r.Db.Model(&orcamento).Association("Arquivos").Replace(foo.Arquivos)
	r.Db.Model(&orcamento).Association("Procedure").Replace(foo.Procedure)
	r.Db.Model(&orcamento).Association("ValorProposta").Replace(foo.ValorProposta)

	err = r.Db.Where("id = ?", id).First(&orcamento).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "Orcamento nao encontrado",
			"error":  err.Error(),
		})
		return err
	}

	err = r.Db.
		Preload("Arquivos").
		Preload("Procedure").
		Preload("ValorProposta").
		Model(&orcamento).
		UpdateColumns(models.Budget{
			Situacao:      foo.Situacao,
			ClienteRefer:  int(customer.ID),
			VendedorRefer: int(vendedor.ID),
		}).Where("id = ?", id).Error

	if err != nil {
		app.Status(http.StatusBadGateway).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(&fiber.Map{
		"message": "Orcamento updated successfully",
		"item":    orcamento,
	})
	return nil
}
