package email

import (
	"AirAccountEmailAdapter/pkg"
	"time"
)

// OpParser get instructions from *op string
func OpParser(msg *string) *pkg.Op {
	return &pkg.Op{
		Timestamp: time.Time{},
		From:      "stub",
		To:        "stub",
		Value:     "stub",
		Message:   *msg,
		OpId:      "stub",
		PrevOpId:  "stub",
	}
}
