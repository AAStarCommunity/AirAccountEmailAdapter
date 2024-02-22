package infra

import (
	"fmt"
	"github.com/knadh/go-pop3"
	"log"
)

func Dial(server string, port int, user, password string) {
	pop3 := pop3.New(pop3.Opt{
		Host:       server,
		Port:       port,
		TLSEnabled: true,
	})
	// Create a new connection. POP3 connections are stateful and should end
	// with a Quit() once the opreations are done.
	c, err := pop3.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit()

	// Authenticate.
	if err := c.Auth(user, password); err != nil {
		log.Fatal(err)
	}

	// Print the total number of messages and their size.
	count, size, _ := c.Stat()
	fmt.Println("total messages=", count, "size=", size)

	// Pull the list of all message IDs and their sizes.
	msgs, _ := c.List(0)
	for _, m := range msgs {
		fmt.Println("id=", m.ID, "size=", m.Size)
	}

	// Pull all messages on the server. Message IDs go from 1 to N.
	for id := 1; id <= count; id++ {
		m, _ := c.Retr(id)

		subject := m.Header.Get("subject")
		sub, _ := decodeRFC2047String(subject)
		fmt.Println(id, "=", sub)

		// To read the multi-part e-mail bodies, see:
		// https://github.com/emersion/go-message/blob/master/example_test.go#L12
	}
}
