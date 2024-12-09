package server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
	redisdb "bigagent_server/db/redis"
	grpc_client "bigagent_server/grpcs/client"
	"bigagent_server/model"
	"bigagent_server/utils/logger"
	responses "bigagent_server/web/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"io/ioutil"
	"net/http"
)

type ServerApi struct{}

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
	req, err := http.NewRequest(c.Request.Method, "http://"+ip+":8010/bigagent/showdata", c.Request.Body)
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
	if global.CONF.System.Serct != c.Request.Header.Get("Authorization") {
		logger.DefaultLogger.Error(fmt.Errorf("配置秘钥认证失败！"))
		responses.FailWithAgent(c, "", "配置秘钥认证失败！")
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "新增config失败！")
		return
	}
	var configDB model.AgentConfigDB
	err = json.Unmarshal(body, &configDB)
	id, err := mysqldb.AgentConfigId()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "新增config失败！")
		return
	}
	configDB.ID = id + 1
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
	if global.CONF.System.Serct != c.Request.Header.Get("Authorization") {
		logger.DefaultLogger.Error(fmt.Errorf("配置秘钥认证失败！"))
		responses.FailWithAgent(c, "", "配置秘钥认证失败！")
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	var requestdata map[string]int
	err = json.Unmarshal(body, &requestdata)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}

	//查询配置
	id, _ := requestdata["config_id"]
	config, err := mysqldb.AgentConfigSelect(id)
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.FailWithAgent(c, "", "查询配置失败")
		return
	}

	responses.SuccssWithDetailed(c, "", "正在下发中，请查看agent状态")

	//从Redis批量获取所有agent地址
	agentAddrs, err := redisdb.ScanAgentAddresses()
	if err != nil {
		logger.DefaultLogger.Error(err)
		//responses.FailWithAgent(c, "", "获取agent地址失败")
		return
	}

	// 如果Redis中没有数据，从MySQL获取并更新Redis
	if len(agentAddrs) == 0 {
		// 从MySQL批量获取并更新Redis
		if err := mysqldb.UpdateAgentAddressesToRedis(); err != nil {
			logger.DefaultLogger.Error(err)
			responses.FailWithAgent(c, "", "更新agent地址缓存失败")
			return
		}
		// 重新从Redis获取
		agentAddrs, err = redisdb.ScanAgentAddresses()
		if err != nil {
			logger.DefaultLogger.Error(err)
			responses.FailWithAgent(c, "", "获取agent地址失败")
			return
		}
	}

	// 发处理配置推送
	errChan := make(chan error, len(agentAddrs))
	for _, addr := range agentAddrs {
		go func(address string) {
			conn, err := grpc_client.InitClient(address)
			if err != nil {
				errChan <- fmt.Errorf("连接agent(%s)失败: %v", address, err)
				return
			}
			defer conn.Close()

			err = grpc_client.GrpcConfigPush(conn, &config, global.CONF.System.Serct)
			if err != nil {
				errChan <- fmt.Errorf("推送配置到agent(%s)失败: %v", address, err)
				return
			}
			errChan <- nil
		}(addr)
	}

	// 收集错误信息
	var failedCount int
	agentUUIDs, _, err := mysqldb.AgentConfigNetSelect(len(agentAddrs))
	for i := 0; i < len(agentUUIDs); i++ {
		if err := <-errChan; err != nil {
			failedCount++
			logger.DefaultLogger.Error(err)
		}
	}

	if failedCount > 0 {
		//responses.SuccssWithAgent(c, "", fmt.Sprintf("配置下发异常，失败%d个", failedCount))
	} else {
		//responses.SuccssWithAgent(c, "", "配置下发成功!")
	}
}

// @Summary 查询Agent配置
// @Description 通过WebSocket实时查询Agent配置
// @Tags Agent配置
// @Accept json
// @Produce json
// @Param Authorization header string true "认证密钥"
// @Router /v1/ws/config [get]
func (*ServerApi) GetAgentConfig(c *gin.Context) {
	// 验证密钥
	if global.CONF.System.Serct != c.Request.Header.Get("Authorization") {
		logger.DefaultLogger.Error(fmt.Errorf("配置秘钥认证失败！"))
		responses.FailWithAgent(c, "", "配置秘钥认证失败！")
		return
	}

	configs, err := mysqldb.AgentConfigSelectAll(c.Query("page"), c.Query("pageSize"))
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.SuccssWithDetailed(c, "", "查询失败！")
		return
	}
	num, err := mysqldb.AgentConfigNetNum()
	if err != nil {
		logger.DefaultLogger.Error(err)
		responses.SuccssWithDetailed(c, "", "查询失败！")
		return
	}
	responses.SuccssWithDetailedFenye(c, "", configs, num)
}
