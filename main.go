package main

import (
	"AirAccountEmailAdapter/conf"
)

func main() {

	c := conf.Get()

	for {
		Quartz(c)
	}
}
