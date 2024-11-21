package inits

import (
	"bigagent_server/utils/logger"
	"bigagent_server/web/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//配置json日志
	r.Use(logger.GinLoggerMiddleware(), logger.GinRecoveryMiddleware())
	// 基础路由
	r1 := router.RouterGroupApp.ServerRouter
	r1.Router(r)
	// 扩展路由
	r2 := router.RouterGroupApp.OtherRouter

	// 数据库读取路径，for循环执行
	for _, otherRouter := range r2 {
		otherRouter.Router("", r)
	}
	return r
}
