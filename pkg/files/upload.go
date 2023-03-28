package files

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Upload(app *fiber.Ctx) error {
	file, err := app.FormFile("document")

	if err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid data",
		})
		return err
	}

	// Save file to root directory:
	err = app.SaveFile(file, fmt.Sprintf("./public/%s", file.Filename))
	if err != nil {
		app.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Upload failed",
		})
		return err
	}

	return app.Status(http.StatusOK).JSON(&fiber.Map{
		"url": "https://dentshow-api.up.railway.app/" + file.Filename,
	})
}
