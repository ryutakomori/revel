package mail

import (
	"log"
	"testing"
)

func Test_Mail(t *testing.T) {
	// send mail
	from := "no-play@test.jp"
	email := "test@gmail.com"
	subject := "ユニットテスト 件名"
	msg := "ユニットテスト 本文"

	err := mail.Send(from, email, subject, msg)
	if err != nil {
		log.Println("send mail error:", err)
	}

}
