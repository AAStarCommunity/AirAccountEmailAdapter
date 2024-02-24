package gateway

import "fmt"

func send() {
	for {
		data := <-ch
		fmt.Println("recv:", data.Message)
	}
}
