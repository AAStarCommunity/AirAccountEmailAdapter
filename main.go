package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
)

func main() {

	c := conf.Get()

	if conn, err := infra.Dial(c.MyEmail.Host, c.MyEmail.TlsPort, c.MyEmail.User, c.MyEmail.Password); err != nil {
		panic(err)
	} else {
		defer conn.Quit()

		infra.Retrieve(conn)
	}
}
