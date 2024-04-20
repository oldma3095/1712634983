package client

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"github.com/oldma3095/1712634983/tools"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Log              *zap.Logger
	Conn             *grpc.ClientConn
	ApiServiceClient commonApi.ApiServiceClient
}

func NewMaster(serverIp string, serverPort int) (*Clients, error) {
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
		Log:              zLog,
		Conn:             conn,
		ApiServiceClient: commonApi.NewApiServiceClient(conn),
	}
	return client, nil
}
