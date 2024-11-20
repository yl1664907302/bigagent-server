package grpc_server

import (
	"context"
	"fmt"
)

type GrpcServer struct {
	UnimplementedPushAgantDataServer
}

func (s GrpcServer) SendData(ctx context.Context, req *StandData) (*ResponseMessage, error) {
	defer func() {
		panicInfo := recover() //panicInfo是any类型，即传给panic()的参数
		if panicInfo != nil {
			fmt.Println(panicInfo)
		}
	}()

	message := ResponseMessage{
		Code:    "200",
		Message: "成功",
	}

	return &message, nil
}
