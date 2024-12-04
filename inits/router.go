package inits

import (
	"bigagent_server/utils/logger"
	"bigagent_server/web/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(Cors)
	//配置json日志
	r.Use(logger.GinLoggerMiddleware(), logger.GinRecoveryMiddleware())
	// 基础路由
	r1 := router.RouterGroupApp.ServerRouter
	r1.Router(r)
	r2 := router.RouterGroupApp.UserRouter
	r2.Router(r)
	// 扩展路由
	r3 := router.RouterGroupApp.OtherRouter
	// 数据库读取路径，for循环执行
	for _, otherRouter := range r3 {
		otherRouter.Router("", r)
	}
	return r
}

// 跨域
func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
	c.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
