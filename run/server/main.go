package main

import "github.com/oldma3095/1712634983/server"

func main() {
	port := 7555
	server.RunGRPCServer(port)
}
