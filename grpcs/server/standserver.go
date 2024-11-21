package grpc_server

import (
	"bigagent_server/db/mysqldb"
	model "bigagent_server/model/agent"
	"bigagent_server/utils/logger"
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"net"
	"time"
)

type GrpcServer struct {
	UnimplementedPushAgantDataServer
}

func (s GrpcServer) SendData(ctx context.Context, req *StandData) (*ResponseMessage, error) {
	// 获取客户端的 IP 地址
	p, ok := peer.FromContext(ctx)
	// 提取 IP 部分（去掉端口）
	host, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, err
	}
	defer func() {
		panicInfo := recover() //panicInfo是any类型，即传给panic()的参数
		if panicInfo != nil {
			fmt.Println(panicInfo)
		}
	}()

	// 查询是否存在agent信息
	_, err = mysqldb.AgentSelect(req.Uuid)
	// 不存在就创建
	if err != nil {
		err = mysqldb.AgentRegister(&model.AgentInfo{
			UUID:         req.Uuid,
			NetIP:        host,
			Hostname:     req.Hostname,
			IPv4First:    req.Ipv4,
			Active:       1,
			Status:       "",
			ActionDetail: "",
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		})
		logger.DefaultLogger.Error(err)
		return &ResponseMessage{
			Code:    "200",
			Message: "agent register success ！",
		}, err
	} else {
		//	存在就执行更新
		err = mysqldb.AgentUpdateAllExceptUUID(req.Uuid, &model.AgentInfo{
			UUID:         req.Uuid,
			NetIP:        host,
			Hostname:     req.Hostname,
			IPv4First:    req.Ipv4,
			Active:       1,
			Status:       "",
			ActionDetail: "",
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		})
		logger.DefaultLogger.Error(err)

	}
	logger.DefaultLogger.Info("ip地址为：", host)
	return &ResponseMessage{
		Code:    "200",
		Message: "agent update success ！",
	}, err
}
