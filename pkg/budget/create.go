package budget

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Create Cria um novo Budget.
//	@Description	Cria um novo Budget
//	@Summary		Cria um novo Budget
//	@Tags			Budget
//	@Accept			json
//	@Produce		json
//	@Param			budget body models.Budget true "Budget"
//	@Success		200	{object} models.Budget
//	@Router			/api/v1/orcamentos [post]
func (r handler) Create(app *fiber.Ctx) error {
	var orcamento models.Budget

	err := app.BodyParser(&orcamento)

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	// errors := models.ValidateStruct(orcamento)
	// if errors != nil {
	// 	return app.Status(fiber.StatusBadRequest).JSON(errors)
	// }

	err = r.Db.Create(&orcamento).Error

	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err,
		})
		return err
	}

	app.JSON(orcamento)
	return nil
}
