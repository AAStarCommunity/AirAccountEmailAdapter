package infra

import (
	"AirAccountEmailAdapter/conf"
	"fmt"
	"net/smtp"
	"strings"
)

const contentType = "Content-Type: text/plain; charset=UTF-8"

func ReplyEmail(replyTo, subject, body string) error {
	mailCfg := conf.Get().Mail
	auth := smtp.PlainAuth("", mailCfg.User, mailCfg.Password, mailCfg.Host)
	to := []string{replyTo}
	replier := mailCfg.Replier
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + replier +
		"<" + mailCfg.User + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(fmt.Sprintf("%s:25", mailCfg.Host), auth, replier, to, msg)
	return err
}
