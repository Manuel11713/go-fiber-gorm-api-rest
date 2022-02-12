package main

import (
	"github.com/Manuel11713/go-auth/database"
	"github.com/Manuel11713/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{ //this way we can receive cookies from frontend
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":5000")

}
