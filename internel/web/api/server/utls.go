package server

import (
	"bigagent_server/internel/logger"
	"bigagent_server/internel/web/response"
	"bigagent_server/internel/web/services"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

func Err(c *gin.Context, err error, keyword string) bool {
	if err != nil {
		logger.DefaultLogger.Error(err) // 记录错误日志

		// 根据错误类型或条件返回不同的响应
		switch keyword {
		case "add": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "添加失败", err)
		case "info": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "查询失败", err)
		case "edit": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "编辑失败", err)
		case "delete": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "删除失败", err)
		case "update": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "更新失败", err)
		case "push": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "配置下发败", err)
		case "host": // 替换为具体的错误类型
			responses.ResponseApp.FailWithAgent(c, "未找到有效主机", err)
		case "sse": // 替换为另一个具体的错误类型
			responses.ResponseApp.FailWithAgentSSE(c, err)
		default:
			responses.ResponseApp.FailWithAgent(c, "请求异常", err) // 默认错误处理
		}
		return true
	}
	return false
}

// 发送sse数据
func sendAgentInfo(c *gin.Context) error {
	// 获取 agent 信息
	info, err := services.AgentServiceImpV1App.GetAgentInfo(c)
	if Err(c, err, "sse") {
		return err
	}

	dnum, anum, err := services.AgentServiceImpV1App.GetAgentNumDead2Live(c)
	if Err(c, err, "sse") {
		return err
	}
	responses.ResponseApp.SuccssWithAgentInfosSSE(c, info, dnum, anum)
	return nil
}

func sendRedict(c *gin.Context, host string, key string) {
	resp, err := services.AgentServiceImpV1App.GetAgentRedictShow(c, host, key, true)
	if Err(c, err, "info") {
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
	c.Writer.Write(bodyBytes)
}
