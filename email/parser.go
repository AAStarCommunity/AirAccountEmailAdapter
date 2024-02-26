package email

import (
	"AirAccountEmailAdapter/pkg"
	"strings"
	"time"
)

type PolicyType int

// OpParser get instructions from *op string
func OpParser(msg *string) *pkg.Op {

	var action pkg.OpActionType
	if strings.EqualFold(*msg, string(pkg.BindWallet)) {
		action = pkg.BindWallet
	} else if strings.EqualFold(*msg, string(pkg.QueryBalance)) {
		action = pkg.QueryBalance
	} else {
		// TODO: the rest action
	}
	return &pkg.Op{
		Action:    action,
		Timestamp: time.Time{},
		From:      "stub",
		To:        "stub",
		Value:     "stub",
		Message:   *msg,
		OpId:      "stub",
		PrevOpId:  "stub",
	}
}
