package gateway

import (
	"AirAccountEmailAdapter/pkg"
)

var ch chan *pkg.Op

func NewMailOp() chan *pkg.Op {
	ch = make(chan *pkg.Op)

	go recv()

	return ch
}
