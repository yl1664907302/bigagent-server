package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccssWithAgent(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func FailWithAgent(c *gin.Context, msg string, data any) {
	// 错误的时候要返回错误code
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"data": data,
	})
}

func SuccssWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func SuccssWithDetailedFenye(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func FailWithDetailed(c *gin.Context, msg string, data any) {
	// 错误的时候要返回错误code
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

// 前端登入需要data中存在code为0
func LoginSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"token":   "123456",
	})
}

func InfoSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	})
}

func GomessageSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"result":  data,
		"error":   "null",
	})
}
