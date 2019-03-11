package mail

import (
	"net/smtp"
	"os"
)

func Send(from string, to string, subject string, msg string) error {

	contents := []byte(
		"To: " + to + "\r\n" +
			"Subject:" + subject + "\r\n" +
			"\r\n" +
			msg)

	err := smtp.SendMail(os.Getenv("MAIL_SERVER_SMTP")+":587",
		smtp.PlainAuth("", os.Getenv("MAIL_SERVER_ACCOUNT"), os.Getenv("MAIL_SERVER_PW"), os.Getenv("MAIL_SERVER_SMTP")),
		from, []string{to}, contents)

	return err
}
