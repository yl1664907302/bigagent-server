package server

import (
	"bigagent_server/config/global"
	"bigagent_server/db/mysqldb"
	model "bigagent_server/model/agentstanddata"
	"bigagent_server/utils/logger"
	responses "bigagent_server/web/response"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	req, err := http.NewRequest(c.Request.Method, "http://"+ip+":"+global.CONF.System.Client_port+"/bigagent/showdata", c.Request.Body)
	if err != nil {
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
		responses.FailWithAgent(c, "查询失败！", err)
		return
	}

	defer resp.Body.Close()

	// 将响应体内容返回给客户端
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	c.Status(resp.StatusCode)

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	c.Writer.Write(bodyBytes)

}

func (*ServerApi) PushAgentData(c *gin.Context) {
	standard := model.NewAgentStandData()
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	err = json.Unmarshal(body, &standard)
	//  匹配密钥并权鉴,
	//  送入指定model构造器，构造数据类型
	//	...
	//  执行reponse
}
