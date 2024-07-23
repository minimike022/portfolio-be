package remail

import (
	"github.com/gofiber/fiber/v2"
	cemail "portfolio/app/controller/email"
)


func SetupEmail(app *fiber.App) {
	app.Get("/send-email", cemail.SendEmail)
}