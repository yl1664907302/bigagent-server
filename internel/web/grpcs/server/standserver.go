package grpc_server

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/db/mysqldb"
	"bigagent_server/internel/model"
	"bigagent_server/internel/utils"
	"context"
	"time"

	"github.com/goccy/go-json"
	"google.golang.org/grpc/metadata"
)

type GrpcServer struct {
	UnimplementedPushAgantDataServer
}

func (s *GrpcServer) SendData(ctx context.Context, req *SmpData) (*ResponseMessage, error) {
	//密钥验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &ResponseMessage{
			Code:    "500",
			Message: "agent serct is error ！",
		}, nil
	}
	tokens := md["authorization"]
	validTokenFound := false
	for _, token := range tokens {
		// 检查每个token是否是期望的Token
		if token == "Bearer "+config.CONF.System.Serct {
			validTokenFound = true
			break
		}
	}
	if !validTokenFound {
		return &ResponseMessage{
			Code:    "500",
			Message: "Authorization token is missing！",
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
