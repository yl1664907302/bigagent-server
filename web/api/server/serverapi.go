package server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
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

func (*ServerApi) PushAgentConfig(c *gin.Context) {
	if global.CONF.System.Serct != c.Request.Header.Get("Authorization") {
		logger.DefaultLogger.Error(fmt.Errorf("配置秘钥认证失败！"))
		responses.FailWithAgent(c, "配置秘钥认证失败！", nil)
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	config := new(model.AgentConfig)
	err = json.Unmarshal(body, config)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	//err = config.UnmarshalAuthDetails(body)
	//if err != nil {
	//	logger.DefaultLogger.Error(err)
	//}
	conn, err := grpc_client.InitClient(global.CONF.System.Grpc_server)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	err = grpc_client.GrpcConfigPush(conn, config, global.CONF.System.Serct)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
}
