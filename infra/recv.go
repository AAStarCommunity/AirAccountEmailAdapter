package infra

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/emersion/go-message"
	"github.com/knadh/go-pop3"
	"github.com/toorop/go-dkim"
	"strings"
	"time"
)

type MailBasis struct {
	MailId  int
	Subject string
	From    string
	To      string
	Date    time.Time
	MsgId   string
}

const patchRFC1123Z = "Mon, 2 Jan 2006 15:04:05 -0700"
const patchRFC1123 = "Mon, 2 Jan 2006 15:04:05 MST"

func Retrieve(conn *pop3.Conn, proc func(basis *MailBasis), mailId int) (int, error) {

	lastMailId := mailId
	msgs, err := conn.List(mailId)
	if err != nil {
		return lastMailId, err
	}

	for _, m := range msgs {
		if msg, err := conn.Retr(m.ID); err == nil {
			if lastMailId < m.ID {
				lastMailId = m.ID
			}
			if verifyDkim(msg) {
				subject := msg.Header.Get("Subject")
				from := msg.Header.Get("From")
				to := msg.Header.Get("To")
				date := func() string {
					pre := msg.Header.Get("Date")
					if i := strings.Index(pre, "("); i > 0 {
						pre = pre[0 : i-1]
					}

					return pre
				}()
				msgId := msg.Header.Get("Message-ID")
				if sub, err := decodeRFC2047String(subject); err == nil {
					go proc(&MailBasis{
						MailId:  m.ID,
						Subject: sub,
						From:    from,
						To:      to,
						Date: func() time.Time {
							loc, _ := time.LoadLocation("UTC")
							if r, err := time.ParseInLocation(time.RFC1123Z, date, loc); err != nil {
								if r, err := time.ParseInLocation(patchRFC1123Z, date, loc); err != nil {
									if r, err := time.ParseInLocation(time.RFC1123, date, loc); err != nil {
										if r, err := time.ParseInLocation(patchRFC1123, date, loc); err != nil {
											return time.Time{}
										} else {
											return r
										}
									} else {
										return r
									}
								} else {
									return r
								}
							} else {
								return r
							}
						}(),
						MsgId: func() string {
							data := []byte(msgId)
							md5Hash := md5.Sum(data)
							return hex.EncodeToString(md5Hash[:])
						}(),
					})
				}
			}
		}
	}

	return lastMailId, nil
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
