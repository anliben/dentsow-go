package users

import (
	"fiber/pkg/common/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *handler) Sign(app *fiber.Ctx) error {
	var user models.User
	userq := new(UserResponse)

	if err := app.BodyParser(userq); err != nil {
		app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "Invalid credentials",
		})
		return nil
	}

	err := r.Db.Where("username = ? and password = ?", userq.Username, userq.Password).First(&user).Error

	if err != nil {
		app.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": err,
		})
		return nil
	}

	exp := time.Now().Add(time.Minute * 30).Unix()

	claims := jwt.MapClaims{
		"name":  user.Username,
		"admin": false,
		"exp":   exp,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return app.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to generate token",
		})
	}

	if err != nil {
		return err
	}

	return app.JSON(fiber.Map{"access": t, "exp": exp, "user": user})
}
