package cemail

import (
	"log"
	memail "portfolio/app/model/email"
	"github.com/gofiber/fiber/v2"
)


func SendEmail(ctx *fiber.Ctx) error {
	email_model := memail.SendEmail{}

	err := ctx.BodyParser(&email_model)
	
	if err != nil {
		log.Fatal(err.Error())
	}

	

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Email Sent Successfully!",
	})
}