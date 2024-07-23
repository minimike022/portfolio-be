package semail


import (
	"gopkg.in/gomail.v2"
)

func Emailtemplate(from string, to string, subject string, body string) *gomail.Message{
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)


	return msg
}