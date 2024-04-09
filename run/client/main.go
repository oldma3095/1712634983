package main

import (
	"github.com/oldma3095/1712634983/client"
	"time"
)

func main() {
	ip := "127.0.0.1"
	port := 7555

	for {
		client.NewMaster(ip, port)
		time.Sleep(time.Second)
	}
}
