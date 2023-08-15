package budget

import (
	"fiber/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Getbudget delete budget.
//	@Description	delete budget.
//	@Summary		delete budget.
//	@Tags			budget
//	@Accept			json
//	@Produce		json
//	@Param			id						path	int		true	"Id"
//	@Param			budget		body	models.Budget	true	"forma pagamento"
//	@Success		200						{array}	models.Budget
//	@Router			/api/v1/orcamentos/{id} [delete]
func (r handler) Delete(app *fiber.Ctx) error {

	id := app.Params("id")

	if id == "" {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID is empty",
		})
		return nil
	}

	resp := r.Db.Raw("DELETE budget_propostas WHERE budget_id = ?", id)

	fmt.Println(resp)

	err := r.Db.Delete(&models.Budget{}, id).Error

	if err != nil {
		err = app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"detail": "Orcamento nao excluido",
			"error":  err.Error(),
		})
		return err
	}

	return app.Status(http.StatusNoContent).JSON(&fiber.Map{})
}
