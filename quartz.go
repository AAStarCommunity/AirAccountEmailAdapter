package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/email"
	"AirAccountEmailAdapter/gateway"
	"AirAccountEmailAdapter/infra"
	"github.com/knadh/go-pop3"
	"time"
)

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

		_ = infra.Retrieve(conn, func(str *string) {
			if op := email.OpParser(str); op != nil {
				ch <- op
			}
		})
	}
}
