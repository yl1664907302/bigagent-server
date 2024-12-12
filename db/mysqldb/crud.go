package mysqldb

import (
	"bigagent_server/config/global"
	redisdb "bigagent_server/db/redis"
	"bigagent_server/model"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func AgentConfigEdit(configID int, newconfig model.AgentConfigDB) error {
	// 检查配置是否存在且未被软删除
	var config model.AgentConfigDB
	err := global.MysqlDataConnect.Unscoped().
		Where("id = ?", configID).
		First(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("配置不存在")
		}
		return fmt.Errorf("查询配置失败: %v", err)
	}

	// 检查是否被软删除
	if !config.DeletedAt.Time.IsZero() {
		return fmt.Errorf("配置已被删除，无法编辑")
	}
	err = global.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", configID).Updates(&newconfig).Error
	return err
}

func AgentConfigDel(configID string) error {
	// 先查询记录是否存在
	var config model.AgentConfigDB
	result := global.MysqlDataConnect.First(&config, configID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %s", configID)
		}
		return fmt.Errorf("查询记录失败: %v", result.Error)
	}

	// 执行软删除
	if err := global.MysqlDataConnect.Delete(&config).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}

	return nil
}

func AgentConfigUpdateTimes(config_id int) error {
	err := global.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", config_id).Update("times", gorm.Expr("times + ?", 1)).Error
	return err
}

func AgentConfigId() (int, error) {
	var maxId int
	err := global.MysqlDataConnect.Model(&model.AgentConfigDB{}).
		Select("COALESCE(MAX(id), 0)").
		Scan(&maxId).Error
	return maxId, err
}

func AgentConfigSelectAll(cp string, ps string) ([]model.AgentConfigDB, error) {
	var agentConfigs []model.AgentConfigDB

	// 将字符串参数转换为整数
	currentPage, err := strconv.Atoi(cp)
	if err != nil {
		return nil, err
	}

	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		return nil, err
	}

	// 计算偏移量
	offset := (currentPage - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询
	err = global.MysqlDataConnect.
		Where("deleted_at IS NULL").
		Limit(pageSize).
		Offset(offset).
		Find(&agentConfigs).Error

	return agentConfigs, err
}

// UpdateAgentAddressesToRedis 从MySQL批量获取并更新到Redis
func UpdateAgentAddressesToRedis() error {
	const (
		batchSize   = 100 // 每批处理的记录数
		workerCount = 5   // 工作协程数量
	)

	// 创建任务通道和错误通道
	taskChan := make(chan map[string]string, workerCount)
	errChan := make(chan error, workerCount)
	doneChan := make(chan bool)

	// 启动工作协程处理Redis写入
	for i := 0; i < workerCount; i++ {
		go func() {
			for batch := range taskChan {
				if err := redisdb.BatchSetAgentAddresses(batch); err != nil {
					errChan <- err
					return
				}
			}
			doneChan <- true
		}()
	}

	// 主协程处理MySQL查询
	go func() {
		var offset int = 0
		for {
			// 分批查询数据
			rows, err := global.MysqlDataConnect.
				Table("agent_info").
				Select("uuid, net_ip", "grpc_port").
				Limit(batchSize).
				Offset(offset).
				Rows()
			if err != nil {
				errChan <- err
				return
			}

			uuidAddressMap := make(map[string]string)
			recordCount := 0

			// 处理当前批次的数据
			for rows.Next() {
				var uuid, ip, port string
				if err := rows.Scan(&uuid, &ip, &port); err != nil {
					rows.Close()
					errChan <- err
					return
				}
				uuidAddressMap[uuid] = ip + ":" + port
				recordCount++
			}
			rows.Close()

			// 如果当前批次有数据，发送到任务通道
			if len(uuidAddressMap) > 0 {
				taskChan <- uuidAddressMap
			}

			// 如果获取的记录数小于批次大小，说明已经处理完所有数据
			if recordCount < batchSize {
				break
			}

			offset += batchSize
		}
		close(taskChan) // 关闭任务通道，通知工作协程没有更多数据
	}()

	// 等待所有工作协程完成或出现错误
	finished := 0
	for {
		select {
		case err := <-errChan:
			return err
		case <-doneChan:
			finished++
			if finished == workerCount {
				return nil
			}
		}
	}
}

func AgentConfigNetNum() (int, error) {
	var num int64
	err := global.MysqlDataConnect.Model(&model.AgentConfigDB{}).Count(&num).Error
	return int(num), err
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
