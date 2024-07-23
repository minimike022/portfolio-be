package memail

type SendEmail struct {
	Full_Name string `json:"full_name"`
	Email string `json:"email"`
	Subject string `json:"subject"`
	Body string `json:"details"`
}