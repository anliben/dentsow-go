package files

import (
	"fiber/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r handler) Upload(app *fiber.Ctx) error {
	form, err := app.MultipartForm()
	var files models.Files

	var arr_file []models.Files

	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {

			new_file_name_uuid := uuid.New().String()

			err = app.SaveFile(fileHeader, fmt.Sprintf("./public/%s", new_file_name_uuid))

			if err != nil {
				err = app.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"detail": "Upload nao realizado! falha ao salvar o arquivo",
					"error":  err,
				})
				return err
			}
			

			erro := r.Db.Create(&models.Files{
				Url:      app.BaseURL() + "/" + new_file_name_uuid,
				Filename: fileHeader.Filename,
			}).Scan(&files)

			arr_file = append(arr_file, files)

			if erro.Error != nil {
				err = app.Status(http.StatusBadRequest).JSON(&fiber.Map{
					"detail": "Upload nao realizado, erro no banco!",
					"error":  erro.Error,
				})
				return err
			}
		}
	}

	return app.Status(http.StatusOK).JSON(&fiber.Map{
		"items": arr_file,
	})
}
