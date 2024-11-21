package mysqldb

import (
	"fmt"
	"testing"
)

func TestAgentSelect(t *testing.T) {
	// 准备测试数据
	testUUID := "test-uuid-12345"
	//testData := model.AgentInfo{
	//	UUID:         testUUID,
	//	NetIP:        "192.168.1.100",
	//	Hostname:     "test-host",
	//	IPv4First:    "192.168.1.1",
	//	Active:       1,
	//	Status:       "running",
	//	ActionDetail: "none",
	//}
	// 检查数据库中是否已删除
	agentSelect, err := AgentSelect(testUUID)
	fmt.Println(agentSelect, err)
}
