package users

import (
	"fiber/pkg/common/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *handler) Sign(app *fiber.Ctx) error {
	var user models.User
	userq := new(UserResponse)

	if err := app.BodyParser(userq); err != nil {
		err = app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "Verifique os dados enviados!",
			"error":  err.Error(),
		})
		return err
	}

	err := r.Db.Where("username = ?", userq.Username).First(&user).Error

	if err != nil {
		err = app.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"detail": "Usuario não encontrado!",
			"error":  err.Error(),
		})
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userq.Password))

	if err != nil {
		err = app.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"detail": "Verifique sua Senha!",
		})
		return err
	}

	exp := time.Now().Add(time.Hour * 5).Unix()

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

	sess, sessErr := store.Get(app)
	if sessErr != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Nenhum refresh identificado" + err.Error(),
		})
	}

	sess.Set(REFRESH, true)
	cookie := sess.Get(REFRESH)

	sessErr = sess.Save()
	if sessErr != nil {
		return app.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Nenhum refresh identificado" + err.Error(),
		})
	}

	return app.JSON(fiber.Map{"access": t, "exp": exp, "user": user, "cookie": cookie})
}
