package test

import (
	"github.com/oldma3095/1712634983/cache"
	"github.com/oldma3095/1712634983/client"
	"github.com/oldma3095/1712634983/server"
	"testing"
	"time"
)

func TestGrpcServer(t *testing.T) {
	uuid := "test"
	port := 7555
	go server.RunGRPCServer(port)

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			info := cache.GetServerSystemInfo(uuid)
			t.Logf("%+v", info)
		}
	}
}

func TestGrpcClient(t *testing.T) {
	ip := "127.0.0.1"
	port := 7555
	uuid := "test"

	c, err := client.NewMaster(ip, port)
	if err != nil {
		t.Fatal(err.Error())
	}
	go c.PushClientInfoToMaster(uuid)
	c.PushResultToMaster(cache.NiuNiuResult{}, nil)
	<-time.After(time.Second * 60)
}
