package db

import "time"

type Mail struct {
	Id          int       `json:"id"`
	Sender      string    `json:"sender"`
	Receiver    string    `json:"receiver"`
	Subject     string    `json:"subject"`
	Unread      bool      `json:"unread"`
	Fingerprint string    `json:"fingerprint"`
	LogAt       time.Time `json:"log_at"`
}
