package email

import (
	"AirAccountEmailAdapter/infra"
	"AirAccountEmailAdapter/pkg"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

type PolicyType int

func Fingerprint(op *pkg.Op) string {
	raw := fmt.Sprintf("%s:%s:%s:%d", op.Action, op.From, op.To, op.Timestamp.Unix())
	data := []byte(raw)
	md5Hash := md5.Sum(data)
	return hex.EncodeToString(md5Hash[:])
}

// OpParser get instructions from *op string
// msg represents basis of the email
func OpParser(mailBasis *infra.MailBasis) *pkg.Op {
	var action pkg.OpActionType
	if strings.EqualFold(mailBasis.Subject, string(pkg.BindWallet)) {
		action = pkg.BindWallet
	} else if strings.EqualFold(mailBasis.Subject, string(pkg.QueryBalance)) {
		action = pkg.QueryBalance
	} else {
		// TODO: the rest action
	}
	return &pkg.Op{
		Action:    action,
		Timestamp: mailBasis.Date,
		From:      mailBasis.From,
		To:        mailBasis.To,
		Message:   mailBasis.Subject,
		OpId:      mailBasis.MsgId,
	}
}
