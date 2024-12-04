package model

import (
	"time"
)

type AgentConfigDB struct {
	ID         int         `gorm:"type:int(11);primary_key;not null;comment:'配置唯一标识符'" json:"id"`
	Title      string      `gorm:"type:varchar(255);not null;comment:'配置标题'" json:"title"`
	Status     string      `gorm:"type:varchar(15);not null;comment:'配置状态'" json:"status"`
	ConfigData AgentConfig `gorm:"type:json;not null;comment:'配置信息'" json:"configdata"`
	Times      string      `gorm:"type:varchar(15);not null;comment:'下发次数'" json:"times"`
	RoleName   string      `gorm:"type:varchar(255);not null;comment:'操作角色'" json:"rolename"`
	Details    string      `gorm:"type:varchar(255);not null;comment:'备注'" json:"details"`
	Ranges     string      `gorm:"type:varchar(255);not null;comment:'操作范围'" json:"ranges"`
	CreatedAt  time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:'创建时间'" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'" json:"updated_at"`
}

func (*AgentConfigDB) TableName() string {
	return "agent_config"
}
