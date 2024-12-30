package responses

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/logger"
	"bigagent_server/internel/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type Response struct {
}

func (r *Response) FailWithAgentSSE(c *gin.Context, err error) {
	response := map[string]any{
		"code":    0,
		"message": err.Error(),
		"data":    "",
	}
	jsonErrorData, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		logger.DefaultLogger.Error(jsonErr)
		return
	}

	// 发送 SSE 格式的错误信息
	fmt.Fprintf(c.Writer, "data: %s\n\n", jsonErrorData)

	// 确保数据立即发送到客户端
	c.Writer.Flush()
}

func (r *Response) SuccssWithAgentInfosSSE(c *gin.Context, agentinfos []model.AgentInfo, dnums int, anums int) {
	// 构造返回数据
	response := map[string]any{
		"code": 0,
		"data": map[string]any{
			"agentInfos": agentinfos,
			"nums":       anums + dnums,
			"dnums":      dnums,
			"anums":      anums,
		},
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return
	}
	// 发送 SSE 格式的数据
	fmt.Fprintf(c.Writer, "data: %s\n\n", jsonData)

	// 确保数据立即发送到客户端
	c.Writer.Flush()
}

func (*Response) SuccssWithAgent(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (*Response) SuccssWithAgentConfigsFenye(c *gin.Context, configs []model.AgentConfigDB, nums int) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]any{
			"configs": configs,
			"nums":    nums,
		},
	})
}

func (*Response) SuccssWithAgentConfigsFail(c *gin.Context, id int, fnum int) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]any{
			"id":   id,
			"fnum": fnum,
		},
	})
}

func (*Response) SuccssWithAgentInfos(c *gin.Context, agentinfos []model.AgentInfo, nums int) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]any{
			"agentinfos": agentinfos,
			"nums":       nums,
		},
	})
}

func (*Response) FailWithAgent(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"data": data,
	})
}

func (*Response) SuccssWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (*Response) SuccssWithDetailedFenye(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (*Response) FailWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (*Response) LoginSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"token":   config.CONF.System.Serct,
	})
}

func (*Response) LoginOutSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"token":   config.CONF.System.Serct,
	})
}

func (*Response) InfoSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	})
}

func (*Response) GomessageSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"result":  data,
		"error":   "null",
	})
}

var ResponseApp = new(Response)
