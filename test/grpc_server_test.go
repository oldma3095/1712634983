package test

import (
	"github.com/oldma3095/1712634983/client"
	"github.com/oldma3095/1712634983/server"
	"testing"
)

func TestGrpcServer(t *testing.T) {
	port := 7555
	server.RunGRPCServer(port)
}

func TestGrpcClient(t *testing.T) {
	ip := "127.0.0.1"
	port := 7555

	client.NewMaster(ip, port)
}
