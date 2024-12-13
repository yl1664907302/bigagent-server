package server

import (
	"bigagent_server/db/mysqldb"
	redisdb "bigagent_server/db/redis"
	grpc_client "bigagent_server/grpcs/client"
	"bigagent_server/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义推送结果结构
type PushResult struct {
	Address string
	Error   error
}

// 定义worker函数
func worker(ctx context.Context, jobs <-chan string, results chan<- PushResult, config *model.AgentConfigDB, secret string) {
	for addr := range jobs {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := grpc_client.InitClient(addr)
			if err != nil {
				results <- PushResult{
					Address: addr,
					Error:   fmt.Errorf("连接agent(%s)失败: %v", addr, err),
				}
				continue
			}
			defer conn.Close()

			err = grpc_client.GrpcConfigPush(conn, config, secret)
			results <- PushResult{
				Address: addr,
				Error:   err,
			}
		}
	}
}

// 获取Agent地址列表(带缓存)
func getAgentAddrs(ctx context.Context, c *gin.Context) ([]string, error) {
	// 先从Redis获取
	addrs, err := redisdb.ScanAgentAddresses(c)
	if err == nil && len(addrs) > 0 {
		return addrs, nil
	}

	// Redis获取失败则从MySQL更新缓存
	if err := mysqldb.UpdateAgentAddressesToRedis(c); err != nil {
		return nil, fmt.Errorf("更新agent地址缓存失败: %v", err)
	}

	// 重新从Redis获取
	return redisdb.ScanAgentAddresses(c)
}
