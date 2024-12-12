package model

import (
	"gorm.io/gorm"
	"time"
)

type AgentConfigDB struct {
	ID        int    `gorm:"type:int(11);primary_key;not null;comment:'配置唯一标识符'" json:"id"`
	Title     string `gorm:"type:varchar(255);not null;comment:'配置标题'" json:"title"`
	Status    string `gorm:"type:varchar(15);not null;comment:'配置状态'" json:"status"`
	Times     string `gorm:"type:varchar(15);not null;comment:'下发次数'" json:"times"`
	RoleName  string `gorm:"type:varchar(255);not null;comment:'操作角色'" json:"role_name"`
	Details   string `gorm:"type:varchar(255);not null;comment:'备注'" json:"details"`
	Ranges    string `gorm:"type:varchar(255);not null;comment:'操作范围'" json:"ranges"`
	Slot_Name string `gorm:"column:solt_name;type:varchar(100);not null" json:"slot_name"`
	AuthName  string `gorm:"type:varchar(255);not null;comment:'授鉴类型'"json:"auth_name"`
	DataName  string `gorm:"type:varchar(255);not null;comment:'数据类型'"json:"data_name"`

	Protocol string `gorm:"type:varchar(255);not null;comment:'网络协议'"json:"protocol"` // 网络协议，如http或https
	Host     string `gorm:"type:varchar(255);not null;comment:'主机信息'"json:"host"`     // 主机地址
	Port     int    `gorm:"type:varchar(255);not null;comment:'端口信息'"json:"port"`     // 端口号
	Path     string `gorm:"type:varchar(255);not null;comment:'路径信息'"json:"path"`     // 请求路径
	//下面为前端传入的各种认证数据
	Token string `json:"token"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updated_at"`
}

type AgentConfigRedis struct {
	Uuid  string `json:"uuid"`
	NetIP string `json:"net_ip"`
}

func (*AgentConfigDB) TableName() string {
	return "agent_config"
}
