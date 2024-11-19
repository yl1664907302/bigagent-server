package inits

import (
	"bigagent_server/config/global"
	grpc_server "bigagent_server/grpcs/server"
	"google.golang.org/grpc"
	"net"
)

func RunG() {
	go func() {
		s := grpc.NewServer()
		server := grpc_server.GrpcServer{}
		grpc_server.RegisterPushAgantDataServer(s, &server)
		// 启动服务
		lis, err := net.Listen("tcp", global.CONF.System.Grpc)
		if err != nil {
			panic(err)
		}
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
}
