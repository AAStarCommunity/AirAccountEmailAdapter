package infra

import (
	"fmt"
	"github.com/knadh/go-pop3"
)

func Retrieve(conn *pop3.Conn) error {

	if count, size, err := conn.Stat(); err != nil {
		return err
	} else {
		fmt.Println("total messages=", count, "size=", size)

		// Pull the list of all message IDs and their sizes.
		msgs, _ := conn.List(0)
		for _, m := range msgs {
			fmt.Println("id=", m.ID, "size=", m.Size)
		}

		// Pull all messages on the server. Message IDs go from 1 to N.
		for id := 1; id <= count; id++ {
			m, _ := conn.Retr(id)

			subject := m.Header.Get("subject")
			sub, _ := decodeRFC2047String(subject)
			fmt.Println(id, "=", sub)

		}
	}

	return nil
}
