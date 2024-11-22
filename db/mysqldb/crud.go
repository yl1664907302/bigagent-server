package mysqldb

import (
	"bigagent_server/config/global"
	model "bigagent_server/model/agent"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func AgentUpdateActiveToDead(t time.Time) ([]model.AgentInfo, error) {
	var agents []model.AgentInfo
	global.MysqlDataConnect.Model(&model.AgentInfo{}).Where("updated_at < ?", t).Find(&agents)
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).
		Where("updated_at < ?", t).
		Update("active", 0).Error
	return agents, err
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
