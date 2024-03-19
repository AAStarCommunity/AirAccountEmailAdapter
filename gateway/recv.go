package gateway

import (
	"AirAccountEmailAdapter/pkg"
	"fmt"
)

func recv() {
	for {
		data := <-ch
		fmt.Println("recv:", data.Message)

		switch data.Action {
		case pkg.BindWallet:
			go bind(data)
		case pkg.QueryBalance:
			go query(data)
		case pkg.Transfer:
			go transfer(data)
		}
	}
}
