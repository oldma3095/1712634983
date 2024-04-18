package client

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/oldma3095/1712634983/cache"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"github.com/oldma3095/1712634983/tools"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Clients struct {
	log              *zap.Logger
	conn             *grpc.ClientConn
	apiServiceClient commonApi.ApiServiceClient
}

func NewMaster(serverIp string, serverPort int) (*Clients, error) {
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				cache.HandleSystemInfo()
			}
		}
	}()
	zLog := tools.Zap()
	address := fmt.Sprintf("%s:%d", serverIp, serverPort)
	//credentials, err := credentials.NewClientTLSFromFile("../pkg/tls/server.pem", "go-grpc-example")
	credentials := insecure.NewCredentials()
	conn, err := grpc.Dial(
		address,
		grpc.WithDefaultCallOptions(grpc.MaxRecvMsgSizeCallOption{
			MaxRecvMsgSize: 512 * 1024 * 1024,
		}),
		grpc.WithTransportCredentials(credentials),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_zap.StreamClientInterceptor(zLog),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_zap.UnaryClientInterceptor(zLog),
		)),
	)
	if err != nil {
		zLog.Error(err.Error())
		return nil, err
	}

	client := &Clients{
		log:              zLog,
		conn:             conn,
		apiServiceClient: commonApi.NewApiServiceClient(conn),
	}
	return client, nil
}
