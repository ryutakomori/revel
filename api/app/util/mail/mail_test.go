package mail

import (
	"log"
	"testing"
)

func Test_Mail(t *testing.T) {

	// send mail

	from := "no-play@test-search.jp"
	email := "test592672@gmail.com"
	subject := "テストサーチ ユニットテスト 件名"
	msg := "テストサーチ ユニットテスト 本文"

	err := mail.Send(from, email, subject, msg)
	if err != nil {
		log.Println("send mail error:", err)
	}

}
