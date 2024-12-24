package model

import (
	"time"
)

type AgentInfo struct {
	UUID         string    `gorm:"column:uuid;primaryKey;type:char(36);not null" json:"uuid"`            // 机器唯一标识符
	NetIP        string    `gorm:"column:net_ip;type:varchar(15);not null" json:"net_ip"`                // 通信IPv4地址
	Hostname     string    `gorm:"column:hostname;type:varchar(255);not null" json:"hostname"`           // 主机名
	IPv4First    string    `gorm:"column:ipv4_first;type:varchar(15);not null" json:"ipv4_first"`        // 首个IPv4地址
	Active       int       `gorm:"column:active;type:tinyint(1);not null;default:1" json:"active"`       // Agent是否在线（1: 是，0: 否）
	Status       string    `gorm:"column:status;type:varchar(255);not null" json:"status"`               // 机器当前状态
	ActionDetail string    `gorm:"column:action_detail;type:varchar(255);not null" json:"action_detail"` // Agent动作描述
	Grpc_port    string    `gorm:"column:grpc_port;type:varchar(100);not null" json:"grpc_port"`
	Platform     string    `gorm:"column:platform;type:varchar(100);not null" json:"platform"`
	Kernel       string    `gorm:"column:kernel;type:varchar(100);not null" json:"kernel"`
	Machine_type string    `gorm:"column:machine_type;type:varchar(100);not null" json:"machine_type"`
	Os           string    `gorm:"column:os;type:varchar(100);not null" json:"os"`
	Arch         string    `gorm:"column:arch;type:varchar(100);not null" json:"arch"`
	Disk_use     string    `gorm:"column:disk_use;type:varchar(100);not null" json:"disk_use"`
	Memory_use   string    `gorm:"column:memory_use;type:varchar(100);not null" json:"memory_use"`
	Cpu_use      string    `gorm:"column:cpu_use;type:varchar(100);not null" json:"cpu_use"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 注册时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName 指定表名，gorm默认使用model方法传入查询表的元信息，并使用"Tabler"接口获取表名
func (a *AgentInfo) TableName() string {
	return "agent_info"
}
