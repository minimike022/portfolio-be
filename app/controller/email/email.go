package cemail

import (
	"fmt"
	"log"
	memail "portfolio/app/model/email"
	semail "portfolio/app/service/email"
	"gopkg.in/gomail.v2"
	"github.com/gofiber/fiber/v2"
)


func SendEmail(ctx *fiber.Ctx) error {
	email_model := memail.SendEmail{}
	
	err := ctx.BodyParser(&email_model)
	fmt.Println(email_model)
	
	if err != nil {
		log.Fatal(err.Error())
	}

	GOMAIL_HOST := `smtp.gmail.com`
	GOMAIL_PORT := 587
	GOMAIL_USER := `solodev.business@gmail.com`
	GOMAIL_PASS := `yjhi yane szed avet`


	receive_email := semail.Emailtemplate(GOMAIL_USER, GOMAIL_USER, email_model.Subject, email_model.Body)
	response_email := semail.Emailtemplate(GOMAIL_USER, email_model.Email, email_model.Subject, "Thank you for your response! We receive your request and will get back to you as soon as we can. Thank you!")
	
	d := gomail.NewDialer(GOMAIL_HOST, GOMAIL_PORT, GOMAIL_USER, GOMAIL_PASS)

	receive := d.DialAndSend(receive_email)
	response := d.DialAndSend(response_email)

	if receive != nil {
		log.Fatal(receive.Error())
	}

	if response != nil {
		log.Fatal(response.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Email Sent Successfully!",
	})
}