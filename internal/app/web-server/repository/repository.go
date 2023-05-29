package repository

type EmailData struct {
	Name    string
	Email   string
	Message string
}

type Repository interface {
	SendMail(recipient string, data *EmailData) error
}
