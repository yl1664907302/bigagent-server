package inits

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/web/grpcs/server"
	"google.golang.org/grpc"
	"net"
)

func RunG() {
	go func() {
		s := grpc.NewServer()

		// 注册服务端
		server := &grpc_server.GrpcServer{}
		grpc_server.RegisterPushAgantDataServer(s, server)

		// 启动服务
		lis, err := net.Listen("tcp", config.CONF.System.Grpc)
		if err != nil {
			panic(err)
		}
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
}
