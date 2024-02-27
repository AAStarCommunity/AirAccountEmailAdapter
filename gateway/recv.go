package gateway

import (
	"AirAccountEmailAdapter/pkg"
	"fmt"
)

func recv() {
	for {
		data := <-ch
		fmt.Println("recv:", data.Message)

		if data.Action == pkg.BindWallet {
			go bind(data)
		}
	}
}
