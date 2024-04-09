package test

import (
	"go_poker/grpc/client"
	"go_poker/grpc/server"
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
