package grpc_server

import (
	"bigagent_server/config"
	"bigagent_server/db/mysqldb"
	"bigagent_server/model"
	"bigagent_server/utils"
	"context"
	"github.com/goccy/go-json"
	"time"
)

type GrpcServer struct {
	UnimplementedPushAgantDataServer
}

func (s *GrpcServer) SendData(ctx context.Context, req *SmpData) (*ResponseMessage, error) {
	//密钥验证
	if config.CONF.System.Serct != req.Serct {
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
	// 查询是否存在agent信息
	_, err = mysqldb.AgentSelect(req.Uuid)
	// 不存在就创建
	if err != nil {
		marshal, _ := json.Marshal(req.DiskUse)
		err = mysqldb.AgentRegister(&model.AgentInfo{
			UUID:         req.Uuid,
			NetIP:        host,
			Hostname:     req.Hostname,
			IPv4First:    req.Ipv4,
			Active:       1,
			Grpc_port:    req.GrpcPort,
			Status:       req.Status,
			Platform:     req.Platform,
			Machine_type: req.MachineType,
			Os:           req.Os,
			Kernel:       req.Kernel,
			Arch:         req.Arch,
			Disk_use:     marshal,
			Cpu_use:      req.CpuUse,
			Memory_use:   req.MemoryUse,
			ActionDetail: req.Actiondetail,
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		})
		return &ResponseMessage{
			Code:    "200",
			Message: "agent register success ！",
		}, err
	} else {
		marshal, _ := json.Marshal(req.DiskUse)
		err = mysqldb.AgentUpdateAllExceptUUID(req.Uuid, &model.AgentInfo{
			UUID:         req.Uuid,
			NetIP:        host,
			Hostname:     req.Hostname,
			IPv4First:    req.Ipv4,
			Active:       1,
			Grpc_port:    req.GrpcPort,
			Status:       req.Status,
			Platform:     req.Platform,
			Machine_type: req.MachineType,
			Os:           req.Os,
			Kernel:       req.Kernel,
			Arch:         req.Arch,
			Disk_use:     marshal,
			Cpu_use:      req.CpuUse,
			Memory_use:   req.MemoryUse,
			ActionDetail: req.Actiondetail,
			UpdatedAt:    time.Time{},
		})
	}
	return &ResponseMessage{
		Code:    "200",
		Message: "agent update success! ",
	}, err
}
