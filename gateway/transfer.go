package gateway

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
	"AirAccountEmailAdapter/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

func transfer(m *pkg.Op) error {
	cfg := conf.Get().AAGateway.Host

	re := regexp.MustCompile(pkg.TransferTo)
	raw := strings.ToLower(m.Message)
	if matches := re.FindStringSubmatch(raw); len(matches) == 3 {
		value := matches[1]
		receiver := matches[2]
		body, _ := json.Marshal(struct {
			Receiver string `json:"receiver"`
			Value    string `json:"value"`
		}{
			Receiver: receiver,
			Value:    value,
		})
		if _, err := http.Post(
			cfg+"/api/instructions/transfer?id="+m.From,
			"application/json",
			bytes.NewBuffer(body)); err != nil {
			return err
		} else {
			b := struct {
				Op string `json:"op"`
			}{}
			go infra.ReplyEmail(m.From, "transfer accepted", m.Message)
			go CheckTransfer(m.From, b.Op)
		}
	}
	return nil
}
