package controllers

import (
	"math/rand"
	"net/smtp"

	"github.com/Manuel11713/go-auth/database"
	"github.com/Manuel11713/go-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := "admin@company.com"

	to := []string{
		data["email"],
	}
	url := "http://localhost:3000/reset/" + token
	message := []byte("click <a href=\"" + url + "\" >here</a> to reset the password")

	err := smtp.SendMail("0.0.0.0:1025", nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords doesn't match!",
		})
	}

	var passwordRest = models.PasswordReset{}

	if err := database.DB.Where("token = ?", data["token"]).Last(&passwordRest); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordRest.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
