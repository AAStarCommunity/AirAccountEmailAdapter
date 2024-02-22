package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
)

func main() {

	c := conf.Get()
	infra.Dial(c.MyEmail.Host, c.MyEmail.TlsPort, c.MyEmail.User, c.MyEmail.Password)
}
