package files

import (
	"fiber/pkg/common/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Delete(app *fiber.Ctx) (err error) {
	var file models.Files

	uuid := app.Params("uuid")

	if uuid == "" {
		err = app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "UUId is empty",
		})
		return err
	}

	err = r.Db.Where("url = ?", app.BaseURL()+"/"+uuid).Delete(&file).Error

	if err != nil {
		err = app.Status(http.StatusNoContent).JSON(&fiber.Map{
			"message": "Media not found",
		})
		return err
	}
	return app.Status(http.StatusOK).JSON(&fiber.Map{})
}
