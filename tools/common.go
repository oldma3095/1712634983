package tools

import (
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"net"
)

func GRPCClientIP(c context.Context) (ip string, err error) {
	fromContext, ok := peer.FromContext(c)
	if !ok {
		err = fmt.Errorf("peer.FromContext fail")
		return
	}
	if tcpAddr, ok2 := fromContext.Addr.(*net.TCPAddr); ok2 {
		ip = tcpAddr.IP.String()
	} else {
		ip = fromContext.Addr.String()
	}
	return
}
