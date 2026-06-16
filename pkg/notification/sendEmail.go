package notification

import (
	"net/smtp"

	"ecommerce-app/config"
)

type NotificationClient interface {
	SendEmail(email string, subject string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func NewNotificationClient(cfg config.AppConfig) NotificationClient {
	return &notificationClient{
		config: cfg,
	}
}

func (c *notificationClient) SendEmail(
	email string,
	subject string,
	message string,
) error {

	msg := []byte(
		"From: " + c.config.SMTPFrom + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			message,
	)

	auth := smtp.PlainAuth(
		"",
		c.config.SMTPUser,
		c.config.SMTPPassword,
		c.config.SMTPHost,
	)

	return smtp.SendMail(
		c.config.SMTPHost+":"+c.config.SMTPPort,
		auth,
		c.config.SMTPFrom,
		[]string{email},
		msg,
	)
}