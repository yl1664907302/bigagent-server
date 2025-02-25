package mysqldb

import (
	conf "bigagent_server/internel/config"
	"bigagent_server/internel/db/redis"
	"bigagent_server/internel/model"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func AgentconfigRangesUpdate(id int, keyword string) error {
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"ranges": keyword}).Error
	return err
}
func AgentconfigRangesBool(id int) (bool, error) {
	// 查询当前的 ranges 值
	var agentConfig model.AgentConfigDB
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", id).First(&agentConfig).Error
	if err != nil {
		return false, err
	}
	if agentConfig.Ranges == "空" || agentConfig.Ranges == "全部" {
		return false, nil
	}
	return true, nil
}

func AgentconfigRangesSelect(id int) ([]string, error) {
	// 查询当前的 ranges 值
	var agentConfig model.AgentConfigDB
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", id).First(&agentConfig).Error
	return strings.Split(agentConfig.Ranges, ","), err
}

func AgentconfigRangesInsert(id int, uuids []string) error {
	// 查询当前的 ranges 值
	var agentConfig model.AgentConfigDB
	if err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", id).First(&agentConfig).Error; err != nil {
		return err
	}

	// 将当前的 ranges 字符串解析为切片
	currentRanges := strings.Split(agentConfig.Ranges, ",")

	// 去重逻辑：只追加不存在的 UUID
	for _, uuid := range uuids {
		// 检查当前 ranges 中是否已经存在该 UUID
		exists := false
		for _, existingUUID := range currentRanges {
			if existingUUID == uuid {
				exists = true
				break
			}
		}

		// 如果不存在，则追加
		if !exists {
			currentRanges = append(currentRanges, uuid)
		}
	}

	// 将切片转换回以逗号分隔的字符串
	updatedRangesStr := strings.Join(currentRanges, ",")

	// 去掉开头的逗号（如果有）
	updatedRangesStr = strings.TrimLeft(updatedRangesStr, ",")

	// 更新数据库中的 ranges 字段
	if err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).
		Where("id = ?", id).
		Update("ranges", updatedRangesStr).Error; err != nil {
		return err
	}

	return nil
}

func AgentconfigStatusChange(id int, s string) error {
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id", id).Update("status", s).Error
	return err
}

func AgentconfigSelectByFail(key string) (int, error) {
	var fnum int64
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Where("action_detail = ?", "config update failed"+"["+key+"]").Count(&fnum).Error
	return int(fnum), err
}

func AgentconfigNewID() (string, error) {
	var id int64
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("deleted_at IS NULL").Select("COALESCE(MAX(id), 0)").Where("times > ?", 0).Scan(&id).Error
	return fmt.Sprintf("%d", id), err
}

func AgentDelete(uuid string) error {
	// 先查询记录是否存在
	err := conf.MysqlDataConnect.Where("uuid = ?", uuid).Delete(&model.AgentInfo{}).Error
	return err
}

func AgentSelectlive2dead() (int, int, error) {
	var anum int64
	var dnum int64
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Where("active = ?", 0).Count(&dnum).Error
	if err != nil {
		return 0, 0, err
	}
	err = conf.MysqlDataConnect.Model(&model.AgentInfo{}).Where("active = ?", 1).Count(&anum).Error
	if err != nil {
		return 0, 0, err
	}
	return int(dnum), int(anum), nil
}

func AgentInfoSelectAll(cp string, ps string) ([]model.AgentInfo, error) {
	var agentInfos []model.AgentInfo
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
	err = conf.MysqlDataConnect.
		Where("deleted_at IS NULL").
		Limit(pageSize).
		Offset(offset).
		Find(&agentInfos).Error
	return agentInfos, err
}

func AgentInfoSelectByKeys(cp string, ps string, uuid string, ip string, t string, p string, a string, c string, c_f string) ([]model.AgentInfo, error) {
	var agentInfos []model.AgentInfo

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

	// 构建查询
	query := conf.MysqlDataConnect.Where("deleted_at IS NULL")

	// 动态添加条件
	if uuid != "" {
		query = query.Where("uuid = ?", uuid)
	}
	if ip != "" {
		query = query.Where("ipv4_first = ?", ip)
	}
	if t != "" {
		query = query.Where("machine_type = ?", t)
	}
	if p != "" {
		query = query.Where("platform = ?", p)
	}
	if a != "" {
		num, _ := strconv.ParseInt(a, 10, 64)
		query = query.Where("active= ?", num)
	}
	if c != "" {
		query = query.Where("action_detail = ?", c)
	}
	if c_f != "" {
		query = query.Where("action_detail LIKE ?", "%failed["+c_f+"]%")
	}
	// 执行查询
	err = query.Limit(pageSize).
		Offset(offset).
		Find(&agentInfos).Error

	return agentInfos, err
}

func AgentNum() (int, error) {
	var num int64
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Count(&num).Error
	return int(num), err
}

func AgentconfigEdit(configID int, newconfig model.AgentConfigDB) error {
	// 检查配置是否存在且未被软删除
	var config model.AgentConfigDB
	err := conf.MysqlDataConnect.Unscoped().
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
	err = conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", configID).Updates(&newconfig).Error
	return err
}

func AgentconfigDel(configID string) error {
	// 先查询记录是否存在
	var config model.AgentConfigDB
	result := conf.MysqlDataConnect.First(&config, configID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %s", configID)
		}
		return fmt.Errorf("查询记录失败: %v", result.Error)
	}

	// 执行软删除
	if err := conf.MysqlDataConnect.Delete(&config).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}

	return nil
}

func AgentconfigUpdateTimes(config_id int) error {
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).Where("id = ?", config_id).Update("times", gorm.Expr("times + ?", 1)).Error
	return err
}

func AgentconfigId() (int, error) {
	var maxId int
	err := conf.MysqlDataConnect.Model(&model.AgentConfigDB{}).
		Select("COALESCE(MAX(id), 0)").
		Scan(&maxId).Error
	return maxId, err
}

func AgentconfigSelectAll(cp string, ps string, status string) ([]model.AgentConfigDB, error) {
	var agentconfigs []model.AgentConfigDB

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

	// 构建查询
	query := conf.MysqlDataConnect.
		Where("deleted_at IS NULL") // 始终添加 deleted_at 条件

	// 如果 status 不为空，添加 status 筛选条件
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 使用 Limit 和 Offset 进行分页查询
	err = query.
		Limit(pageSize).
		Offset(offset).
		Find(&agentconfigs).Error

	return agentconfigs, err
}

// UpdateAgentAddressesToRedis 从MySQL批量获取并更新到Redis
func UpdateAgentAddressesToRedis(c context.Context) error {
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
				if err := redisdb.BatchSetAgentAddresses(c, batch); err != nil {
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
			rows, err := conf.MysqlDataConnect.
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

func AgentconfigNetNum(status string) (int, error) {
	var num int64
	query := conf.MysqlDataConnect.Model(&model.AgentConfigDB{})

	// 如果 status 不为空，添加 WHERE 条件
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 统计记录数量
	err := query.Count(&num).Error
	return int(num), err
}

func AgentconfigNetSelect(num int) ([]string, []string, error) {
	var agents []model.AgentInfo
	// 使用 Select 查询多个字段（net_ip 和 uuid）
	err := conf.MysqlDataConnect.
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

func AgentconfigSelect(id int) (model.AgentConfigDB, error) {
	var agentconfigDB model.AgentConfigDB
	err := conf.MysqlDataConnect.Model(model.AgentConfigDB{}).Where("id = ?", id).First(&agentconfigDB).Error
	return agentconfigDB, err
}

func AgentconfigCreate(c model.AgentConfigDB) error {
	err := conf.MysqlDataConnect.Create(&c).Error
	return err
}

func LoginUser(username string, password string) (model.User, error) {
	var user model.User
	err := conf.MysqlDataConnect.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err
}

func AgentNetIPSelectByUuid(uuid string) (string, error) {
	var agent model.AgentInfo
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Select("net_ip").Where("uuid = ?", uuid).First(&agent).Error
	if err != nil {
		return "", err
	}
	return agent.NetIP, nil
}

func FindDeadAgents(t time.Time) ([]model.AgentInfo, error) {
	var agents []model.AgentInfo
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Where("updated_at < ?", t).Find(&agents).Error
	return agents, err
}

func UpdateDeadAgents(t time.Time) error {
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).Where("updated_at < ?", t).Omit("updated_at").Update("active", 0).Update("status", "部分数据推送异常").Error
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
	err := conf.MysqlDataConnect.Create(&a).Error
	err = redisdb.SetAgentAddresses(context.Background(), a.UUID, a.NetIP+":"+a.Grpc_port)
	return err
}

func AgentUpdateAllExceptUUID(uuid string, a *model.AgentInfo) error {
	err := conf.MysqlDataConnect.Model(&model.AgentInfo{}).
		Where("uuid = ?", uuid).
		Omit("uuid").Omit("created_at").
		Updates(a).Error
	err = redisdb.SetAgentAddresses(context.Background(), a.UUID, a.NetIP+":"+a.Grpc_port)
	return err
}

func AgentSelect(uuid string) (*model.AgentInfo, error) {
	var a model.AgentInfo
	err := conf.MysqlDataConnect.Where("uuid = ?", uuid).First(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("no record found with uuid: %s", uuid)
	}

	return &a, err
}
