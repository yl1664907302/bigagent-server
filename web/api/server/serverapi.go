package server

import (
	model "bigagent_server/model/agentstanddata"
	"bigagent_server/utils/logger"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ServerApi struct{}

func (*ServerApi) SearchAgent(c *gin.Context) {

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
