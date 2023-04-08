package users

import "github.com/gofiber/fiber/v2"

func (r handler) Logout(app *fiber.Ctx) error {

	sess, err := store.Get(app)

	if err != nil {
		return app.Status(fiber.StatusOK).JSON(fiber.Map{
			"detail": "logged out (no session)",
		})
	}

	err = sess.Destroy()
	if err != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "something went wrong: " + err.Error(),
		})
	}

	return app.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}
