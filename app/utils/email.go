package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
)

var authentication smtp.Auth

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

// NewRequest ...
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

// RegisterNotification ...
func RegisterNotification(to, userName, loginName string) {
	password := os.Getenv("DEFAULT_PASSWORD_USER")
	url := os.Getenv("APP_WEB")

	templateData := struct {
		Name, Username, Password, URL string
	}{
		Name:     userName,
		Username: loginName,
		Password: password,
		URL:      url,
	}

	r := NewRequest([]string{to}, "Register SAS", "Welcome to SAS")

	err := r.ParseTemplate("register.html", templateData)
	if err == nil {
		_, errSend := r.SendEmail()
		if errSend != nil {
			fmt.Println(errSend, "ini ya")
		}
	} else {
		log.Println(err)
	}

}

// SendEmail ...
func (r *Request) SendEmail() (bool, error) {

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPass := os.Getenv("SMTP_PASS")

	authentication = smtp.PlainAuth("", smtpEmail, smtpPass, smtpHost)

	mime := "MIME-version: 1.0;\nContent-Type: text/html;"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	if err := smtp.SendMail(addr, authentication, "sasnoreply66@gmail.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

// ParseTemplate ...
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	filePrefix, _ := filepath.Abs("./app/utils/")
	t, err := template.ParseFiles(filePrefix + "/" + templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
