package email

import (
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

const (
	kFrom         = "Peng Renliang <pengrenliang_robot@163.com>"
	kFromUser     = "pengrenliang_robot@163.com"
	kFromPassword = "rghgmvobgqnhbqzh"
	kHost         = "smtp.163.com"
)

func Send(tos []string, subject string, body string) error {
	e := email.Email{
		To:      tos,
		From:    kFrom,
		Subject: subject,
		HTML:    []byte(body),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(
		"smtp.163.com:25",
		smtp.PlainAuth(
			"",
			kFromUser,
			kFromPassword,
			kHost,
		),
	)
	return err
}
