package grpc_client

import (
	"bigagent_server/logger"
	"bigagent_server/model"
	"bigagent_server/web/grpcs/grpc_config"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"time"
)

// InitClient 通用grpc客户端
func InitClient(host string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond) //连接超时设置为1000毫秒
	defer cancel()
	//连接到服务端
	conn, err := grpc.DialContext(
		ctx,
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()), //Credential即使为空，也必须设置
		grpc.WithBlock(), //grpc.WithBlock()直到连接真正建立才会返回，否则连接是异步建立的。因此grpc.WithBlock()和Timeout结合使用才有意义。server端正常的情况下使用grpc.WithBlock()得到的connection.GetState()为READY，不使用grpc.WithBlock()得到的connection.GetState()为IDEL
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(10<<20), grpc.MaxCallRecvMsgSize(10<<20)), //默认情况下SendMsg上限是MaxInt32，RecvMsg上限是4M，这里都修改为10M
	)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func GrpcConfigPush(conn *grpc.ClientConn, config *model.AgentConfigDB, serct string) error {
	client := grpc_config.NewAgentConfigServiceClient(conn)
	//准备好请求参数
	agentConfig := grpc_config.AgentConfig{
		Id:       strconv.Itoa(config.ID),
		Serct:    serct,
		AuthName: config.AuthName,
		DataName: config.DataName,
		Token:    config.Token,
		SlotName: config.Slot_Name,
		NetworkInfo: &grpc_config.NetworkInfo{
			Protocol: config.Protocol,
			Host:     config.Host,
			Port:     int64(config.Port),
			Path:     config.Path,
		},
	}
	response, err := client.PushAgentConfig(context.Background(), &agentConfig)
	if err != nil {
		logger.DefaultLogger.Error("推送配置失败:", err)
		return err
	} else {
		logger.DefaultLogger.Infof("配置推送成功：%s", response)
	}
	return err
}
