package users

import "github.com/gofiber/fiber/v2"

func (r handler) HealthCheck(app *fiber.Ctx) error {
	sess, err := store.Get(app)

	if err != nil {
		return app.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "not authorized",
		})
	}

	auth := sess.Get(REFRESH)

	if auth != nil {
		return app.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "authenticated",
		})
	} else {
		return app.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "not authorized",
		})
	}
}
