package server

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	commonApi "github.com/oldma3095/1712634983/protos/common/api"
	"github.com/oldma3095/1712634983/server/api"
	"github.com/oldma3095/1712634983/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func RunGRPCServer(port int) {
	zLog := tools.Zap()
	// 监听本地端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 创建gRPC服务器
	s := grpc.NewServer(
		grpc.MaxSendMsgSize(512*1024*1024),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zLog),
			//grpc_auth.StreamServerInterceptor(middleware.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zLog),
			//grpc_auth.UnaryServerInterceptor(myAuthFunction),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			//MaxConnectionIdle:     15 * time.Second, // 如果客户端空闲 15 秒，发送 GOAWAY
			//MaxConnectionAge:      30 * time.Second, // 如果任何连接存在超过 30 秒，发送一个 GOAWAY
			//MaxConnectionAgeGrace: 5 * time.Second,  // 在强行关闭连接之前，等待 5 秒等待挂起的 RPC 完成
			Time:    3 * time.Second, // 如果客户端空闲 3 秒，请 Ping 客户端以确保连接仍处于活动状态
			Timeout: 1 * time.Second, // 等待 1 秒等待 ping 确认，然后假设连接已断开
		}),
	)
	// 注册服务
	{
		commonApi.RegisterApiServiceServer(s, &api.CommonServer{}) // common
	}
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
