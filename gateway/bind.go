package gateway

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
	"AirAccountEmailAdapter/pkg"
	"bytes"
	"net/http"
)

func bind(m *pkg.Op) error {
	cfg := conf.Get().AAGateway.Host
	if resp, err := http.Post(cfg+"/api/instructions/bind?id="+m.From, "application/json", bytes.NewBuffer([]byte("{}"))); err != nil {
		return err
	} else {
		if resp.StatusCode == http.StatusOK {
			go infra.ReplyEmail(m.From, "Congratulations! Your AirAccount Created!", m.Message)
		} else if resp.StatusCode == http.StatusNotAcceptable {
			go infra.ReplyEmail(m.From, "Your AirAccount Already Exists!", m.Message)
		} else {
			return err
		}
	}
	return nil
}
