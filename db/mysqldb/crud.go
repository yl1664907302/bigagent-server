package mysqldb

import (
	"bigagent_server/config/global"
	model "bigagent_server/model/agent"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

//func LoginUser(username string, password string) (pojo.User, error) {
//	var user pojo.User
//	//err := global.MysqlDataConnect.Select("username,password").Find(&p).Error
//	err := global.MysqlDataConnect.Where("username = ? AND password = ?", username, password).First(&user).Error
//	log.Println(user)
//	return user, err
//}

func AgentRegister(a *model.AgentInfo) error {
	err := global.MysqlDataConnect.Create(&a).Error
	return err
}

func AgentUpdateAllExceptUUID(uuid string, a *model.AgentInfo) error {
	// 使用 Omit 排除 uuid 字段
	err := global.MysqlDataConnect.Model(&model.AgentInfo{}).
		Where("uuid = ?", uuid).
		Omit("uuid").
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
