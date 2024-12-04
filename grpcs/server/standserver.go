package grpc_server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
	"bigagent_server/model"
	"bigagent_server/utils"
	"bigagent_server/utils/logger"
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
			Status:       "online",
			ActionDetail: "register",
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		})
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
			Status:       "online",
			ActionDetail: "register",
			UpdatedAt:    time.Time{},
		})
	}
	logger.DefaultLogger.Info("ip地址为：", host)
	logger.DefaultLogger.Infof("agent信息为：%s", req)
	return &ResponseMessage{
		Code:    "200",
		Message: "agent update success! ",
	}, err
}
