package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

// 如果不需要检查重复，可以使用 SetNX (SET if Not eXists)
func TestSetAgentAddresses_test(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	key := AgentAddressPrefix + "03560274-043C-050D-8C06-800700080009"
	client.SetEX(context.Background(), key, "10.0.0.1", AgentAddressTTL).Err()
	return
}
