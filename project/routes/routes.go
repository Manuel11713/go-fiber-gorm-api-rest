package routes

import (
	"github.com/Manuel11713/go-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)

	app.Post("/register", controllers.Regisger)
	app.Post("/login", controllers.Login)

	app.Get("/user", controllers.User)

	app.Post("/logout", controllers.LogOut)
	app.Post("/forgot", controllers.Forgot)
	app.Post("/reset", controllers.Reset)

}
