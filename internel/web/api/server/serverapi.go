package server

import (
	conf "bigagent_server/internel/config"
	"bigagent_server/internel/db/redis"
	"bigagent_server/internel/logger"
	"bigagent_server/internel/model"
	"bigagent_server/internel/web/grpcs/client"
	"bigagent_server/internel/web/response"
	"bigagent_server/internel/web/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type ServerApi struct{}

// GetAgentConfigFail @Summary 搜索最近发布的的Agent配置中失败数量
// @Description 搜索离线Agent数量
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info [get]
func (*ServerApi) GetAgentConfigFail(c *gin.Context) {
	id, fnum, err := services.AgentServiceImpV1App.GetAgentConfigNEW2Fail(c)
	if Err(c, err, "info") {
		return
	}
	responses.ResponseApp.SuccssWithAgentConfigsFail(c, id, fnum)
}

// GetAgentNumDead @Summary 搜索离线Agent数量
// @Description 搜索离线Agent数量
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info [get]
func (*ServerApi) GetAgentNumDead(c *gin.Context) {
	dnum, _, err := services.AgentServiceImpV1App.GetAgentNumDead2Live(c)
	if Err(c, err, "info") {
		return
	}
	responses.ResponseApp.SuccssWithAgent(c, "", dnum)
}

// DeleteAgentInfo  @Summary 删除Agent
// @Description 查询Agent的基本信息
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info [get]
func (*ServerApi) DeleteAgentInfo(c *gin.Context) {
	err := services.AgentServiceImpV1App.DeleteAgentInfo(c)
	if Err(c, err, "delete") {
		return
	}
	responses.ResponseApp.SuccssWithAgent(c, "", "删除成功")
}

// GetAgentInfo @Summary 搜索Agent
// @Description 查询Agent的基本信息
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info [get]
func (*ServerApi) GetAgentInfo(c *gin.Context) {
	info, err := services.AgentServiceImpV1App.GetAgentInfo(c)
	if Err(c, err, "info") {
		return
	}
	num, err := services.AgentServiceImpV1App.GetAgentNum(c)
	if Err(c, err, "info") {
		return
	}
	responses.ResponseApp.SuccssWithAgentInfos(c, info, num)
}

// GetAgentInfoSSE @Summary sse协议分页查询Agent
// @Description sse协议分页查询Agent的基本信息
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info_sse [get]
func (*ServerApi) GetAgentInfoSSE(c *gin.Context) {
	// 设置 SSE 相关的 headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 创建一个 ticker 定期发送数据
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 创建一个 channel 用于检测客户端断开连接
	clientGone := c.Writer.CloseNotify()

	// 首次加载时发送数据
	if err := sendAgentInfo(c); err != nil {
		log.Println(err)
		return
	}

	for {
		select {
		case <-clientGone:
			return
		case <-ticker.C:
			// 定期发送数据
			if err := sendAgentInfo(c); err != nil {
				return
			}
		}
	}
}

// SearchAgent @Summary 搜索Agent
// @Description 根据UUID查询Agent并转发请求
// @Tags Agent管理
// @Accept json
// @Produce json
// @Param uuid query string true "Agent UUID"
// @Router /v1/agent_id[get]
func (*ServerApi) SearchAgent(c *gin.Context) {
	ip, err := services.AgentServiceImpV1App.SearchAgentNet(c)
	if Err(c, err, "info") {
		return
	}
	sendRedict(c, ip+conf.CONF.System.Agent_port, "showdata")
}

// SearchAgentPatrol  @Summary 巡检Agent
// @Description 根据UUID查询Agent巡检数据
// @Tags Agent管理
// @Accept json
// @Produce json
// @Param uuid query string true "Agent UUID"
// @Router /v1/agent_patrol[get]
func (*ServerApi) SearchAgentPatrol(c *gin.Context) {
	ip, err := services.AgentServiceImpV1App.SearchAgentNet(c)
	if Err(c, err, "info") {
		return
	}
	sendRedict(c, ip+conf.CONF.System.Agent_port, "patroldata")
}

// AddAgentConfig @Summary 添加Agent配置
// @Description 新增Agent的配置信息
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Param config body model.AgentConfigDB true "Agent配置信息"
// @Router /v1/add [post]
func (*ServerApi) AddAgentConfig(c *gin.Context) {
	err := services.AgentServiceImpV1App.AddAgentConfig(c)
	if Err(c, err, "add") {
		return
	}
	responses.ResponseApp.SuccssWithAgent(c, "", "添加成功！")
}

// PushAgentConfig @Summary 推送Agent配置
// @Description 向所有在线Agent推送指定配置
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Param config_id body int true "配置ID"
// @Router /v1/push [post]
func (*ServerApi) PushAgentConfig(c *gin.Context) {
	var agentAddrs []string
	var status string
	var config *model.AgentConfigDB
	// 读取请求体并缓存
	body, err := c.GetRawData()
	if Err(c, err, "push") {
		return
	}
	responses.ResponseApp.SuccssWithDetailed(c, "正在操作配置中，请查看agent状态", "正在操作配置中，请查看agent状态")

	//判断配置文件是否存在指定范围
	co, hosts, s, key, err := services.AgentServiceImpV1App.GetAgentConfig2Nets(c, body)
	if key {
		co2, uuids, s2, err2 := services.AgentServiceImpV1App.GetAgentConfig2Uuids(c, body, true)
		if err2 != nil {
			if Err(c, err2, "push") {
				return
			}
		}
		// 验证uuid是否有效，获取有效主机
		validHosts := make([]string, 0)
		for _, uuid := range uuids {
			if exists, host := redisdb.CheckAgentExists(c, uuid); exists {
				validHosts = append(validHosts, host)
			}
		}
		agentAddrs = validHosts
		status = s2
		config = co2
	} else {
		config = co
		agentAddrs = hosts
		status = s
	}

	// 并发推送配置
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results := make(chan error, len(agentAddrs))
	semaphore := make(chan struct{}, 10) // 限制并发数为10
	for _, addr := range agentAddrs {
		semaphore <- struct{}{} // 获取信号量
		go func(address string) {
			defer func() { <-semaphore }() // 释放信号量

			conn, err := grpc_client.InitClient(address, conf.CONF.System.Serct)
			if err != nil {
				results <- fmt.Errorf("连接agent(%s)失败: %v", address, err)
				return
			}
			defer conn.Close()

			if err := grpc_client.GrpcConfigPush(conn, config, conf.CONF.System.Serct); err != nil {
				results <- fmt.Errorf("推送到agent(%s)失败: %v", address, err)
				return
			}
			results <- nil
		}(addr)
	}

	// 收集结果
	for i := 0; i < len(agentAddrs); i++ {
		select {
		case err = <-results:
			if err != nil {
				if Err(c, err, "push") {
					return
				}
			}
		case <-ctx.Done():
			//特殊处理
			responses.ResponseApp.FailWithAgent(c, "", "配置操作超时")
			return
		}
	}
	err = services.AgentServiceImpV1App.UpdateAgentConfigStatus(c, config.ID, status)
	if Err(c, err, "update") {
		return
	}
}

// PushAgentConfigByHost @Summary 下发指定主机的Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/push_host [post]
func (*ServerApi) PushAgentConfigByHost(c *gin.Context) {
	// 读取请求体并缓存
	body, err := c.GetRawData()
	if Err(c, err, "push") {
		return
	}
	config, uuids, status, err := services.AgentServiceImpV1App.GetAgentConfig2Uuids(c, body, false)
	if Err(c, err, "push") {
		return
	}
	// 验证uuid是否有效，获取有效主机
	validHosts := make([]string, 0)
	for _, uuid := range uuids {
		if exists, host := redisdb.CheckAgentExists(c, uuid); exists {
			validHosts = append(validHosts, host)
		}
	}
	if len(validHosts) == 0 && Err(c, fmt.Errorf(""), "host") {
		return
	}
	responses.ResponseApp.SuccssWithDetailed(c, "", "指定agent，正在更新配置")

	// 并发推送配置
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results := make(chan error, len(validHosts))
	semaphore := make(chan struct{}, 5) // 限制并发数为5

	for _, addr := range validHosts {
		semaphore <- struct{}{}
		go func(address string) {
			defer func() { <-semaphore }()

			conn, err := grpc_client.InitClient(address, conf.CONF.System.Serct)
			if err != nil {
				results <- fmt.Errorf("连接agent(%s)失败: %v", address, err)
				return
			}
			defer conn.Close()

			if err = grpc_client.GrpcConfigPush(conn, config, conf.CONF.System.Serct); err != nil {
				results <- fmt.Errorf("推送到agent(%s)失败: %v", address, err)
				return
			}
			results <- nil
		}(addr)
	}

	// 收集结果
	for i := 0; i < len(validHosts); i++ {
		select {
		case err := <-results:
			if err != nil {
				logger.DefaultLogger.Error(err)
			}
		case <-ctx.Done():
			// 特殊处理
			responses.ResponseApp.FailWithAgent(c, "", "配置推送超时")
			return
		}
	}

	// 异步更新配置状态
	go func() {
		err := services.AgentServiceImpV1App.UpdateAgentConfigStatus(c, config.ID, status)
		if Err(c, err, "update") {
			return
		}
	}()
}

// GetAgentConfig @Summary 查询Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/get [get]
func (*ServerApi) GetAgentConfig(c *gin.Context) {
	configs, num, err := services.AgentServiceImpV1App.GetAgentConfigs2num(c)
	if Err(c, err, "info") {
		return
	}
	responses.ResponseApp.SuccssWithAgentConfigsFenye(c, configs, num)
}

// DelAgentConfig @Summary 删除Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/del [delete]
func (*ServerApi) DelAgentConfig(c *gin.Context) {
	err := services.AgentServiceImpV1App.DelAgentConfig(c)
	if Err(c, err, "delete") {
		return
	}
	responses.ResponseApp.SuccssWithDetailed(c, "", "配置删除成功！")
	return
}

// EditAgentConfig @Summary 修改Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/edit [put]
func (*ServerApi) EditAgentConfig(c *gin.Context) {
	err := services.AgentServiceImpV1App.EditAgentConfig(c)
	if Err(c, err, "edit") {
		return
	}
	responses.ResponseApp.SuccssWithAgent(c, "", "编辑config成功！")
}
