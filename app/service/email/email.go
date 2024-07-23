package semail


import (
	"gopkg.in/gomail.v2"
	"bytes"
	"html/template"

	memail "portfolio/app/model/email"
)

func Emailtemplate(from string, to string, subject string, body string) *gomail.Message{
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)


	return msg
}

func SendEmail(data memail.SendEmail, renderedTemplate string, email string) error {
	GOMAIL_HOST := `smtp.gmail.com`
	GOMAIL_PORT := 587
	GOMAIL_USER := `solodev.business@gmail.com`
	GOMAIL_PASS := `yjhi yane szed avet`

	msg := EmailTemplate(GOMAIL_USER, email, data.Subject, renderedTemplate)

	dialer := gomail.NewDialer(GOMAIL_HOST, GOMAIL_PORT, GOMAIL_USER, GOMAIL_PASS)

	app_msg := dialer.DialAndSend(msg)

	return app_msg
}

func EmailTemplate(from string, to string, subject string, body string) *gomail.Message{
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	return msg
}

func RenderTemplate(templateFile string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer
	err = tmpl.Execute(&renderedTemplate, data)
	if err != nil {
		return "", err
	}

	return renderedTemplate.String(), nil
}