package test

import (
	"github.com/oldma3095/1712634983/client"
	"github.com/oldma3095/1712634983/server"
	"testing"
	"time"
)

func TestGrpcServer(t *testing.T) {
	port := 7555
	server.RunGRPCServer(port)
}

func TestGrpcClient(t *testing.T) {
	ip := "127.0.0.1"
	port := 7555

	c, err := client.NewMaster(ip, port)
	if err != nil {
		t.Fatal(err.Error())
	}
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				c.PushClientInfoToMaster()
			}
		}
	}()
	c.PushResultToMaster()
}
