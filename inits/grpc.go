package inits

import (
	"bigagent_server/config"
	grpc_server2 "bigagent_server/web/grpcs/server"
	"google.golang.org/grpc"
	"net"
)

func RunG() {
	go func() {
		s := grpc.NewServer()

		// 注册服务端
		server := &grpc_server2.GrpcServer{}
		grpc_server2.RegisterPushAgantDataServer(s, server)

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
