package pkg

import "time"

type Op struct {
	Timestamp time.Time
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Message   string `json:"message"`
	OpId      string `json:"opId"`     // op hashed value, unique global
	PrevOpId  string `json:"prevOpId"` // represents the op which trigger this op produced if not empty.
}
