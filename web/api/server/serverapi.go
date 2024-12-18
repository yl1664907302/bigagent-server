package server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
	redisdb "bigagent_server/db/redis"
	grpc_client "bigagent_server/grpcs/client"
	"bigagent_server/model"
	"bigagent_server/utils/logger"
	responses "bigagent_server/web/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ServerApi struct{}

// @Summary 搜索Agent
// @Description 查询Agent的基本信息
// @Tags Agent管理
// @Accept json
// @Produce json
// @Router /v1/info [get]
func (*ServerApi) GetAgentInfo(c *gin.Context) {
	agentInfos, err := mysqldb.AgentInfoSelectAll(c.Query("page"), c.Query("pageSize"))
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "agent信息查询失败！", err)
		return
	}
	num, err := mysqldb.AgentNum()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "agent信息查询失败！", err)
		return
	}
	responses.SuccssWithDetailedFenye(c, "", map[string]any{
		"agentInfos": agentInfos,
		"nums":       num,
	})
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func (*ServerApi) GetAgentInfoWS(c *gin.Context) {
	// 将 HTTP 连接升级为 WebSocket 连接
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return
	}
	defer ws.Close()
	// 创建定时器，定期发送数据
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 获取 agent 信息
			agentInfos, err := mysqldb.AgentInfoSelectAll(c.Query("page"), c.Query("pageSize"))
			if err != nil {
				logger.DefaultLogger.Error(err)
				// 发送错误消息
				ws.WriteJSON(map[string]any{
					"code": 500,
					"msg":  "agent信息查询失败",
				})
				return
			}
			num, err := mysqldb.AgentNum()
			if err != nil {
				logger.DefaultLogger.Error(err)
				ws.WriteJSON(map[string]any{
					"code": 500,
					"msg":  "agent数量查询失败",
				})
				return
			}
			// 发送数据
			err = ws.WriteJSON(map[string]any{
				"code": 200,
				"data": map[string]any{
					"agentInfos": agentInfos,
					"nums":       num,
				},
			})

			if err != nil {
				logger.DefaultLogger.Error(err)
				return
			}
		}
	}
}

// @Summary 搜索Agent
// @Description 根据UUID查询Agent并转发请求
// @Tags Agent管理
// @Accept json
// @Produce json
// @Param uuid query string true "Agent UUID"
// @Router /bigagent/showdata [get]
func (*ServerApi) SearchAgent(c *gin.Context) {
	ip, err := mysqldb.AgentNetIPSelectByUuid(c.Query("uuid"))
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "查询失败！", err)
		return
	}
	// 创建一个新的http.Client实例
	client := &http.Client{}
	// 创建一个新的请求对象，复制原始请求信息
	req, err := http.NewRequest(c.Request.Method, "http://"+ip+":8010/"+c.Query("model_name")+"/showdata", c.Request.Body)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "查询失败！", err)
		return
	}
	// 复制所有请求头
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Add("Authorization", global.CONF.System.Serct)
	// 此处密钥配置会在前端编写完后替换，改为从server的uri获取
	//req.Header.Add("Authorization", c.Request.Header.Get("Authorization"))

	// 发送请求并获取响应
	resp, err := client.Do(req)

	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "查询失败！", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.DefaultLogger.Error(err)
		}
	}(resp.Body)

	// 将响应体内容返回给客户端
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	c.Status(resp.StatusCode)

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	_, err = c.Writer.Write(bodyBytes)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
}

// @Summary 添加Agent配置
// @Description 新增Agent的配置信息
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Param config body model.AgentConfigDB true "Agent配置信息"
// @Router /v1/add [post]
func (*ServerApi) AddAgentConfig(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "新增config失败！")
		return
	}
	var configDB model.AgentConfigDB
	err = json.Unmarshal(body, &configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "新增config失败！")
		return
	}
	configDB.Status = "有效"
	err = mysqldb.AgentConfigCreate(configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "新增config失败！")
		return
	}
	responses.SuccssWithAgent(c, "", "新增config成功！")
}

// @Summary 推送Agent配置
// @Description 向所有在线Agent推送指定配置
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Param config_id body int true "配置ID"
// @Router /v1/push [post]
func (*ServerApi) PushAgentConfig(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 1. 获取配置ID并查询配置
	var requestdata map[string]int
	if body, err := c.GetRawData(); err != nil {
		responses.FailWithAgent(c, "", "获取请求数据失败")
		return
	} else if err = json.Unmarshal(body, &requestdata); err != nil {
		responses.FailWithAgent(c, "", "解析请求数据失败")
		return
	}

	id := requestdata["config_id"]
	config, err := mysqldb.AgentConfigSelect(id)
	if err != nil {
		responses.FailWithAgent(c, "", "查询配置失败")
		return
	}

	// 2. 获取agent地址列表
	agentAddrs, err := redisdb.ScanAgentAddresses(c)
	if err != nil || len(agentAddrs) == 0 {
		if err := mysqldb.UpdateAgentAddressesToRedis(c); err != nil {
			responses.FailWithAgent(c, "", "更新agent地址失败")
			return
		}
		agentAddrs, _ = redisdb.ScanAgentAddresses(c)
	}

	responses.SuccssWithDetailed(c, "", "正在下发中，请查看agent状态")

	// 3. 并发推送配置
	results := make(chan error, len(agentAddrs))
	semaphore := make(chan struct{}, 10) // 限制并发数为10

	for _, addr := range agentAddrs {
		semaphore <- struct{}{} // 获取信号量
		go func(address string) {
			defer func() { <-semaphore }() // 释放信号量

			conn, err := grpc_client.InitClient(address)
			if err != nil {
				results <- fmt.Errorf("连接agent(%s)失败: %v", address, err)
				return
			}
			defer conn.Close()

			if err := grpc_client.GrpcConfigPush(conn, &config, global.CONF.System.Serct); err != nil {
				results <- fmt.Errorf("推送到agent(%s)失败: %v", address, err)
				return
			}
			results <- nil
		}(addr)
	}

	// 4. 收集结果
	var failedCount int
	for i := 0; i < len(agentAddrs); i++ {
		select {
		case err := <-results:
			if err != nil {
				failedCount++
				logger.DefaultLogger.Error(err)
			}
		case <-ctx.Done():
			responses.FailWithAgent(c, "", "配置推送超时")
			return
		}
	}

	// 5. 异步更新配置使用次数
	go func() {
		err := mysqldb.AgentConfigUpdateTimes(id)
		if err != nil {
			logger.DefaultLogger.Error(err)
		}
	}()
}

// @Summary 下发指定主机的Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/push_host [post]
func (*ServerApi) PushAgentConfigByHost(c *gin.Context) {
	// 1. 获取请求数据
	var requestData struct {
		ConfigID int      `json:"config_id"`
		Uuids    []string `json:"uuids"` // 主机IP列表
	}
	if body, err := c.GetRawData(); err != nil {
		responses.FailWithAgent(c, "", "获取请求数据失败")
		return
	} else if err = json.Unmarshal(body, &requestData); err != nil {
		responses.FailWithAgent(c, "", "解析请求数据失败")
		return
	}

	// 查询配置信息
	config, err := mysqldb.AgentConfigSelect(requestData.ConfigID)
	if err != nil {
		responses.FailWithAgent(c, "", "查询配置失败")
		return
	}

	// 验证uuid是否有效
	validHosts := make([]string, 0)
	for _, uuid := range requestData.Uuids {
		if exists, host := redisdb.CheckAgentExists(c, uuid); exists {
			validHosts = append(validHosts, host)
		}
	}

	if len(validHosts) == 0 {
		responses.FailWithAgent(c, "", "未找到有效的目标主机")
		return
	}

	responses.SuccssWithDetailed(c, "", "指定agent，正在更新配置")

	// 4. 并发推送配置
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results := make(chan error, len(validHosts))
	semaphore := make(chan struct{}, 5) // 限制并发数为5

	for _, addr := range validHosts {
		semaphore <- struct{}{}
		go func(address string) {
			defer func() { <-semaphore }()

			conn, err := grpc_client.InitClient(address)
			if err != nil {
				results <- fmt.Errorf("连接agent(%s)失败: %v", address, err)
				return
			}
			defer conn.Close()

			if err := grpc_client.GrpcConfigPush(conn, &config, global.CONF.System.Serct); err != nil {
				results <- fmt.Errorf("推送到agent(%s)失败: %v", address, err)
				return
			}
			results <- nil
		}(addr)
	}

	// 5. 收集结果
	var failedCount int
	for i := 0; i < len(validHosts); i++ {
		select {
		case err := <-results:
			if err != nil {
				failedCount++
				logger.DefaultLogger.Error(err)
			}
		case <-ctx.Done():
			responses.FailWithAgent(c, "", "配置推送超时")
			return
		}
	}

	// 6. 异步更新配置使用次数
	go func() {
		if err := mysqldb.AgentConfigUpdateTimes(requestData.ConfigID); err != nil {
			logger.DefaultLogger.Error(err)
		}
	}()
}

// @Summary 查询Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/get [get]
func (*ServerApi) GetAgentConfig(c *gin.Context) {
	configs, err := mysqldb.AgentConfigSelectAll(c.Query("page"), c.Query("pageSize"))
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.SuccssWithDetailed(c, "", "查询失败！")
		return
	}
	num, err := mysqldb.AgentConfigNetNum()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "查询失败！")
		return
	}
	responses.SuccssWithDetailedFenye(c, "", map[string]any{
		"configs": configs,
		"nums":    num,
	})
}

// @Summary 删除Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/del [delete]
func (*ServerApi) DelAgentConfig(c *gin.Context) {
	id := c.Param("config_id")
	err := mysqldb.AgentConfigDel(id)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "删除失败！")
		return
	}
	responses.SuccssWithDetailed(c, "", "删除成功！")
	return
}

// @Summary 修改Agent配置
// @Description
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/edit [put]
func (*ServerApi) EditAgentConfig(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "编辑config失败:"+err.Error())
		return
	}
	var configDB model.AgentConfigDB
	err = json.Unmarshal(body, &configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "编辑config失败:"+err.Error())
	}
	err = mysqldb.AgentConfigEdit(configDB.ID, configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "编辑config失败:"+err.Error())
		return
	}
	responses.SuccssWithAgent(c, "", "编辑config成功！")
}
