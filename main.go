package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	remail "portfolio/app/router/email"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	remail.SetupEmail(app)
	app.Listen(":3000")
}