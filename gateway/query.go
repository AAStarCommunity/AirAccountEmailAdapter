package gateway

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
	"AirAccountEmailAdapter/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Qb struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Balance string `json:"balance"`
		Unit    string `json:"unit"`
		Status  int    `json:"status"`
	} `json:"data"`
	Cost string `json:"cost"`
}

func query(m *pkg.Op) error {
	cfg := conf.Get().AAGateway.Host
	if resp, err := http.Get(cfg + "/api/instructions/balance?id=" + m.From); err != nil {
		return err
	} else {
		if data, err := io.ReadAll(resp.Body); err == nil {
			b := Qb{}
			if err := json.Unmarshal(data, &b); err != nil {
				return err
			}
			go infra.ReplyEmail(m.From, fmt.Sprintf("Your balance is %s %s", b.Data.Balance, b.Data.Unit), m.Message)
		}
	}
	return nil
}
