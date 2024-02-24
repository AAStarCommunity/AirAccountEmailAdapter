package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/email"
	"AirAccountEmailAdapter/infra"
	"github.com/knadh/go-pop3"
)

func main() {

	c := conf.Get()

	if conn, err := infra.Dial(c.MyEmail.Host, c.MyEmail.TlsPort, c.MyEmail.User, c.MyEmail.Password); err != nil {
		panic(err)
	} else {
		defer func(conn *pop3.Conn) {
			_ = conn.Quit()
		}(conn)

		_ = infra.Retrieve(conn, func(str *string) {
			_ = email.OpParser(str)
		})
	}
}
