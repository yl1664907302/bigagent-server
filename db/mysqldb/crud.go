package mysqldb

import (
	"bigagent_server/config/global"
	"bigagent_server/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func AgentConfigNetSelect() ([]string, error) {
	var netIPs []string // 定义一个切片来存储查询结果

	// 使用 Pluck 仅查询 "net_ip" 字段
	err := global.MysqlDataConnect.
		Model(&model.AgentInfo{}).
		Pluck("net_ip", &netIPs).
		Error
	if err != nil {
		return nil, err
	}
	return netIPs, err
}

func AgentConfigSelect(id int) (model.AgentConfig, error) {
	var agentConfigDB model.AgentConfigDB
	err := global.MysqlDataConnect.Model(model.AgentConfigDB{}).Where("id = ?", id).First(&agentConfigDB).Error
	return agentConfigDB.ConfigData, err
}

func AgentConfigCreate(c model.AgentConfigDB) error {
	err := global.MysqlDataConnect.Create(&c).Error
	return err
}

func LoginUser(username string, password string) (model.User, error) {
	var user model.User
	err := global.MysqlDataConnect.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}

func AgentNetIPSelectByUuid(uuid string) (string, error) {
	var agent model.AgentInfo
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).Select("net_ip").Where("uuid = ?", uuid).First(&agent).Error
	if err != nil {
		return "", err
	}
	return agent.NetIP, nil
}

func FindDeadAgents(t time.Time) ([]model.AgentInfo, error) {
	var agents []model.AgentInfo
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).Where("updated_at < ?", t).Find(&agents).Error
	return agents, err
}

func UpdateDeadAgents(t time.Time) error {
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).Where("updated_at < ?", t).Omit("updated_at").Update("active", 0).Error
	return err
}

func AgentUpdateActiveToDead(t time.Time) ([]model.AgentInfo, error) {
	agents, err := FindDeadAgents(t)
	if err != nil {
		return nil, err
	}
	err = UpdateDeadAgents(t)
	if err != nil {
		return nil, err
	}
	return agents, nil
}
func AgentRegister(a *model.AgentInfo) error {
	err := global.MysqlDataConnect.Create(&a).Error
	return err
}

func AgentUpdateAllExceptUUID(uuid string, a *model.AgentInfo) error {
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).
		Where("uuid = ?", uuid).
		Omit("uuid").Omit("created_at").
		Updates(a).Error
	return err
}

func AgentSelect(uuid string) (*model.AgentInfo, error) {
	var a model.AgentInfo
	err := global.MysqlDataConnect.Where("uuid = ?", uuid).First(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("no record found with uuid: %s", uuid)
	}

	return &a, err
}
