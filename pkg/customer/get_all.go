package customer

import (
	"fiber/pkg/common/models"
	"fiber/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @title Fiber Swagger Example API
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} []models.Customer
// @Router / [get]
func (r handler) GetAll(app *fiber.Ctx) error {
	var customer []models.Customer

	cliente_nome := app.Query("cliente_nome")
	cliente_data_nascimento := app.Query("cliente_data_nascimento")
	cliente_cpf := app.Query("cliente_cpf")
	cliente_rg := app.Query("cliente_rg")
	cliente_email := app.Query("cliente_email")
	cliente_idade := app.Query("cliente_idade")
	cliente_estado_civil := app.Query("cliente_estado_civil")
	cliente_sexo := app.Query("cliente_sexo")
	cliente_contato := app.Query("cliente_contato")
	cliente_contato2 := app.Query("cliente_contato2")
	cliente_cep := app.Query("cliente_cep")
	cliente_logradouro := app.Query("cliente_logradouro")
	cliente_numero := app.Query("cliente_numero")
	cliente_complemento := app.Query("cliente_complemento")
	cliente_bairro := app.Query("cliente_bairro")
	cliente_cidade := app.Query("cliente_cidade")
	cliente_estado := app.Query("cliente_estado")
	cliente_prontuario := app.Query("cliente_prontuario")
	cliente_situacao := app.Query("cliente_situacao")
	cliente_indicacao := app.Query("cliente_indicacao")
	cliente_profissao := app.Query("cliente_profissao")
	cliente_observacao := app.Query("cliente_observacao")
	cliente_consultas_creditos := app.Query("cliente_consultas_creditos")
	cliente_consultas_realizadas := app.Query("cliente_consultas_realizadas")
	cliente_consultas_restantes := app.Query("cliente_consultas_restantes")
	cliente_nome_responsavel := app.Query("cliente_nome_responsavel")
	cliente_cpf_responsavel := app.Query("cliente_cpf_responsavel")

	db := r.Db.
		Session(&gorm.Session{PrepareStmt: true}).
		Preload("Midia")

	db.Where(utils.Builder("nome LIKE ?", "%"+cliente_nome+"%"))
	db.Where(utils.Builder("data_nascimento LIKE ?", "%"+cliente_data_nascimento+"%"))
	db.Where(utils.Builder("cpf LIKE ?", "%"+cliente_cpf+"%"))
	db.Where(utils.Builder("rg LIKE ?", "%"+cliente_rg+"%"))
	db.Where(utils.Builder("email LIKE ?", "%"+cliente_email+"%"))
	db.Where(utils.Builder("idade LIKE ?", "%"+cliente_idade+"%"))
	db.Where(utils.Builder("estado_civil LIKE ?", "%"+cliente_estado_civil+"%"))
	db.Where(utils.Builder("sexo LIKE ?", "%"+cliente_sexo+"%"))
	db.Where(utils.Builder("contato LIKE ?", "%"+cliente_contato+"%"))
	db.Where(utils.Builder("contato2 LIKE ?", "%"+cliente_contato2+"%"))
	db.Where(utils.Builder("cep LIKE ?", "%"+cliente_cep+"%"))
	db.Where(utils.Builder("logradouro LIKE ?", "%"+cliente_logradouro+"%"))
	db.Where(utils.Builder("numero LIKE ?", "%"+cliente_numero+"%"))
	db.Where(utils.Builder("complemento LIKE ?", "%"+cliente_complemento+"%"))
	db.Where(utils.Builder("bairro LIKE ?", "%"+cliente_bairro+"%"))
	db.Where(utils.Builder("cidade LIKE ?", "%"+cliente_cidade+"%"))
	db.Where(utils.Builder("estado LIKE ?", "%"+cliente_estado+"%"))
	db.Where(utils.Builder("prontuario LIKE ?", "%"+cliente_prontuario+"%"))
	db.Where(utils.Builder("situacao LIKE ?", "%"+cliente_situacao+"%"))
	db.Where(utils.Builder("indicacao LIKE ?", "%"+cliente_indicacao+"%"))
	db.Where(utils.Builder("profissao LIKE ?", "%"+cliente_profissao+"%"))
	db.Where(utils.Builder("observacao LIKE ?", "%"+cliente_observacao+"%"))
	db.Where(utils.Builder("consultas_creditos LIKE ?", "%"+cliente_consultas_creditos+"%"))
	db.Where(utils.Builder("consultas_realizadas LIKE ?", "%"+cliente_consultas_realizadas+"%"))
	db.Where(utils.Builder("consultas_restantes LIKE ?", "%"+cliente_consultas_restantes+"%"))
	db.Where(utils.Builder("nome_responsavel LIKE ?", "%"+cliente_nome_responsavel+"%"))
	db.Where(utils.Builder("cpf_responsavel LIKE ?", "%"+cliente_cpf_responsavel+"%"))

	err := db.Find(&customer).Error

	if err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	return app.JSON(&fiber.Map{
		"count":    len(customer),
		"next":     "null",
		"previous": "null",
		"items":    customer,
	})
}
