package redisdb

import (
	"bigagent_server/internel/config"
	"context"
	"fmt"
	"time"
)

const (
	AgentAddressPrefix = "agent:addr:"
	AgentAddressTTL    = 1 * time.Hour
	ScanBatchSize      = 100 // 每次扫描的批次大小
)

// ScanAgentAddresses 使用SCAN命令批量获取所有agent地址
func ScanAgentAddresses(c context.Context) ([]string, error) {
	var addresses []string
	pattern := AgentAddressPrefix + "*"
	cursor := uint64(0)

	for {
		var keys []string
		var err error
		// 使用SCAN命令批量扫描key
		keys, cursor, err = config.RedisDataConnect.Scan(c, cursor, pattern, ScanBatchSize).Result()
		if err != nil {
			return nil, err
		}

		// 如果有找到key，批量获取它们的值
		if len(keys) > 0 {
			// 使用MGET批量获取值
			values, err := config.RedisDataConnect.MGet(c, keys...).Result()
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

		// ���果cursor为0，说明已经扫描完所有key
		if cursor == 0 {
			break
		}
	}

	return addresses, nil
}

// BatchSetAgentAddresses 批量设置agent地址
func BatchSetAgentAddresses(c context.Context, uuidAddressMap map[string]string) error {
	pipe := config.RedisDataConnect.Pipeline()

	for uuid, addr := range uuidAddressMap {
		key := AgentAddressPrefix + uuid
		pipe.Set(c, key, addr, AgentAddressTTL)
	}

	_, err := pipe.Exec(c)
	return err
}

// 如果不需要检查重复，可以使用 SetNX (SET if Not eXists)
func SetAgentAddresses(c context.Context, uuid string, addr string) error {
	key := AgentAddressPrefix + uuid
	err := config.RedisDataConnect.SetEX(c, key, addr, AgentAddressTTL).Err()
	return err
}

func UpdateAgentAddress(c context.Context, uuid string, addr string) error {
	key := AgentAddressPrefix + uuid

	// 1. 先检查 key 是否存在
	exists, err := config.RedisDataConnect.Exists(c, key).Result()
	if err != nil {
		return fmt.Errorf("check key exists error: %w", err)
	}

	if exists == 0 {
		return fmt.Errorf("key %s not exists", key)
	}

	// 2. 更新已存在的 key 的值
	err = config.RedisDataConnect.Set(c, key, addr, AgentAddressTTL).Err()
	if err != nil {
		return fmt.Errorf("update key error: %w", err)
	}

	return nil
}

// GetAgentNum 获取agent数量
func GetAgentNum(c context.Context) (int, error) {
	pattern := AgentAddressPrefix + "*"
	cursor := uint64(0)
	count := 0
	for {
		var keys []string
		var err error
		keys, cursor, err = config.RedisDataConnect.Scan(c, cursor, pattern, ScanBatchSize).Result()
		if err != nil {
			return 0, err
		}

		count += len(keys)

		// 如果cursor为0，说明已经扫描完所有key
		if cursor == 0 {
			break
		}
	}
	return count, nil
}

func CheckAgentExists(ctx context.Context, key string) (bool, string) {
	// 使用Get命令获取key的值
	val, err := config.RedisDataConnect.Get(ctx, AgentAddressPrefix+key).Result()
	if err != nil {
		// 如果key不存在或发生错误，返回false和空字符串
		return false, ""
	}
	// 如果key存在，返回true和对应的值
	return true, val
}
