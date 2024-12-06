package redisdb

import (
	"bigagent_server/config/global"
	"context"
	"time"
)

const (
	AgentAddressPrefix = "agent:addr:"
	AgentAddressTTL    = 1 * time.Hour
	ScanBatchSize      = 100 // 每次扫描的批次大小
)

// ScanAgentAddresses 使用SCAN命令批量获取所有agent地址
func ScanAgentAddresses() ([]string, error) {
	var addresses []string
	ctx := context.Background()
	pattern := AgentAddressPrefix + "*"
	cursor := uint64(0)

	for {
		var keys []string
		var err error
		// 使用SCAN命令批量扫描key
		keys, cursor, err = global.RedisDataConnect.Scan(ctx, cursor, pattern, ScanBatchSize).Result()
		if err != nil {
			return nil, err
		}

		// 如果有找到key，批量获取它们的值
		if len(keys) > 0 {
			// 使用MGET批量获取值
			values, err := global.RedisDataConnect.MGet(ctx, keys...).Result()
			if err != nil {
				return nil, err
			}

			// 将非空值添加到结果集
			for _, value := range values {
				if value != nil {
					addresses = append(addresses, value.(string))
				}
			}
		}

		// 如果cursor为0，说明已经扫描完所有key
		if cursor == 0 {
			break
		}
	}

	return addresses, nil
}

// BatchSetAgentAddresses 批量设置agent地址
func BatchSetAgentAddresses(uuidAddressMap map[string]string) error {
	ctx := context.Background()
	pipe := global.RedisDataConnect.Pipeline()

	for uuid, addr := range uuidAddressMap {
		key := AgentAddressPrefix + uuid
		pipe.Set(ctx, key, addr, AgentAddressTTL)
	}

	_, err := pipe.Exec(ctx)
	return err
}

func SetAgentAddresses(uuid string, addr string) error {
	ctx := context.Background()
	key := AgentAddressPrefix + uuid
	err := global.RedisDataConnect.Set(ctx, key, addr, AgentAddressTTL).Err()
	return err

}
