package main

import "go_poker/grpc/server"

func main() {
	port := 7555
	server.RunGRPCServer(port)
}
