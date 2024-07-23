package cemail

import (
	"fmt"
	"log"
	memail "portfolio/app/model/email"
	semail "portfolio/app/service/email"
	"github.com/gofiber/fiber/v2"
)


func SendEmail(ctx *fiber.Ctx) error {
	email_model := memail.SendEmail{}
	
	err := ctx.BodyParser(&email_model)
	fmt.Println(email_model)
	
	if err != nil {
		log.Fatal(err.Error())
	}

	data := struct {
		Full_Name string
		Email string
		Subject string
		Body string

	}{
		Full_Name:            email_model.Full_Name,
		Email:           email_model.Email,
		Subject: email_model.Subject,
		Body: email_model.Body,
	}

	response_path := `C:/Users/Ace Gates/Documents/Portfolio/portfolio-be/app/controller/email/response.html`
	response_template, _ := semail.RenderTemplate(response_path, data)
	receive_path := `C:/Users/Ace Gates/Documents/Portfolio/portfolio-be/app/controller/email/receive.html`
	receive_template, _ := semail.RenderTemplate(receive_path, data)

	receive_email := semail.SendEmail(email_model, receive_template, `solodev.business@gmail.com`)
	if receive_email != nil {
		log.Fatal(receive_email.Error())
	}

	response_email := semail.SendEmail(email_model, response_template, data.Email)
	if response_email != nil {
		log.Fatal(response_email.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Email Sent Successfully!",
	})
}