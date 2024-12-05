package mysqldb

import (
	"bigagent_server/config/global"
	"bigagent_server/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func AgentConfigNetNum() (int64, error) {
	var num int64
	err := global.MysqlDataConnect.Model(&model.AgentConfigDB{}).Count(&num).Error
	return num, err
}

func AgentConfigNetSelect(num int) ([]string, []string, error) {
	var agents []model.AgentInfo
	// 使用 Select 查询多个字段（net_ip 和 uuid）
	err := global.MysqlDataConnect.
		Model(&model.AgentInfo{}).
		Select("net_ip", "uuid"). // 查询 net_ip 和 uuid 字段
		Limit(num).
		Find(&agents).Error
	if err != nil {
		return nil, nil, err
	}

	// 创建一个切片存储 net_ip 字段的值
	var netIPs []string
	for _, agent := range agents {
		netIPs = append(netIPs, agent.NetIP) // 假设 AgentInfo 中有 NetIP 字段
	}
	var uuids []string
	for _, agent := range agents {
		uuids = append(uuids, agent.UUID)
	}
	return uuids, netIPs, nil
}

func AgentConfigSelect(id int) (model.AgentConfigDB, error) {
	var agentConfigDB model.AgentConfigDB
	err := global.MysqlDataConnect.Model(model.AgentConfigDB{}).Where("id = ?", id).First(&agentConfigDB).Error
	return agentConfigDB, err
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
