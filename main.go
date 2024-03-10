package main

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/email/repository/migrations"
)

func main() {

	c := conf.Get()

	migrations.AutoMigrate()

	for {
		Quartz(c)
	}
}
