package gateway

import (
	"AirAccountEmailAdapter/conf"
	"AirAccountEmailAdapter/infra"
	"fmt"
	"net/http"
	"time"
)

func CheckTransfer(from string, op string) {
	cfg := conf.Get().AAGateway.Host

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 10)

		if resp, err := http.Get(cfg + "/api/instructions/transfer/check?id=" + from + "&op=" + op); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("transfer check result:" + resp.Status)

			if resp.StatusCode == 200 {
				go infra.ReplyEmail(from, "transfer successful", "")
			}
		}
	}
}
