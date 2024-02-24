package infra

import (
	"bytes"
	"github.com/emersion/go-message"
	"github.com/knadh/go-pop3"
	"github.com/toorop/go-dkim"
)

func Retrieve(conn *pop3.Conn, proc func(sub *string)) error {

	msgs, err := conn.List(0)
	if err != nil {
		return err
	}

	for _, m := range msgs {
		if msg, err := conn.Retr(m.ID); err == nil {
			if verifyDkim(msg) {
				subject := msg.Header.Get("subject")
				if sub, err := decodeRFC2047String(subject); err == nil {
					proc(&sub)
				}
			}
		}
	}

	return nil
}

func verifyDkim(msg *message.Entity) bool {
	buffer := bytes.NewBuffer(nil)
	if msg.WriteTo(buffer) == nil {
		b := buffer.Bytes()
		if status, err := dkim.Verify(&b); err == nil {
			return status == dkim.SUCCESS || status == dkim.TESTINGSUCCESS
		}
	}
	return false
}
