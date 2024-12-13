package grpc_server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
	redisdb "bigagent_server/db/redis"
	"bigagent_server/model"
	"bigagent_server/utils"
	"context"
	"fmt"
	"time"
)

type GrpcServer struct {
	UnimplementedPushAgantDataServer
}

func (s GrpcServer) SendData(ctx context.Context, req *SmpData) (*ResponseMessage, error) {
	//密钥验证
	if global.CONF.System.Serct != req.Serct {
		return &ResponseMessage{
			Code:    "200",
			Message: "agent serct is error ！",
		}, nil
	}
	// 获取客户端的 IP 地址
	host, err := utils.GetIPToCtx(ctx)
	if err != nil {
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
			Grpc_port:    req.GrpcPort,
			Status:       req.Status,
			ActionDetail: req.Actiondetail,
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		})
		go redisdb.SetAgentAddresses(ctx, req.Uuid, host).Error()
		return &ResponseMessage{
			Code:    "200",
			Message: "agent register success ！",
		}, err
	} else {
		err = mysqldb.AgentUpdateAllExceptUUID(req.Uuid, &model.AgentInfo{
			UUID:         req.Uuid,
			NetIP:        host,
			Hostname:     req.Hostname,
			IPv4First:    req.Ipv4,
			Active:       1,
			Grpc_port:    req.GrpcPort,
			Status:       req.Status,
			ActionDetail: req.Actiondetail,
			UpdatedAt:    time.Time{},
		})
	}
	return &ResponseMessage{
		Code:    "200",
		Message: "agent update success! ",
	}, err
}
