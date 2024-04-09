package client

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	commonApi "go_poker/grpc/protos/common/api"
	"go_poker/grpc/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	log              *zap.Logger
	exit             chan string
	conn             *grpc.ClientConn
	apiServiceClient commonApi.ApiServiceClient
}

func NewMaster(serverIp string, serverPort int) {
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
		return
	}

	client := &Clients{
		log:              zLog,
		exit:             make(chan string),
		conn:             conn,
		apiServiceClient: commonApi.NewApiServiceClient(conn),
	}
	go client.PushClientInfoToMaster()

	exitMsg := <-client.exit
	zLog.Error(exitMsg)
	return
}
