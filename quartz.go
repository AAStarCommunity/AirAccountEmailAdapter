package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/email"
	"AirAccountEmailAdapter/email/repository"
	"AirAccountEmailAdapter/gateway"
	"AirAccountEmailAdapter/infra"
	"github.com/knadh/go-pop3"
	"time"
)

var mailId = 0

func init() {
	db := conf.GetDB()

	db.Model(&repository.Mail{}).Select("MAX(mailId)").Scan(&mailId)
}

func Quartz(c *conf.Conf) {
	timer := time.NewTimer(time.Second * 5)

	<-timer.C

	if conn, err := infra.Dial(c.Mail.Host, c.Mail.Port, c.Mail.User, c.Mail.Password); err != nil {
		panic(err)
	} else {
		defer func(conn *pop3.Conn) {
			_ = conn.Quit()
		}(conn)

		ch := gateway.NewMailOp()

		mailId, _ = infra.Retrieve(conn, func(basis *infra.MailBasis) {
			if op := email.OpParser(basis); op != nil {
				fp := email.Fingerprint(op)
				if err = repository.Save(&repository.Mail{
					MailId:      basis.MailId,
					Sender:      op.From,
					SendAt:      op.Timestamp,
					Receiver:    op.To,
					Subject:     op.Message,
					Unread:      false,
					Fingerprint: fp,
				}); err == nil {
					ch <- op
				}
			}
		}, func() int {
			if mailId == 0 {
				return 0
			} else {
				return mailId + 1
			}
		}())
	}
}
