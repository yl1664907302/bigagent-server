package inits

import (
	"bigagent_server/config"
	"bigagent_server/logger"
	responses "bigagent_server/web/response"
	"bigagent_server/web/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Router() *gin.Engine {
	r := gin.Default()
	//配置json日志
	r.Use(logger.GinRecoveryMiddleware(), AuthMiddleware(), InitCORS())
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

func InitCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 添加 OPTIONS
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Content-Type",
			"Accept",
			"X-Requested-With",
			"Access-Control-Allow-Methods",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// AuthMiddleware 验证密钥中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// 验证密钥
		if config.CONF.System.Serct != c.Request.Header.Get("Authorization") {
			logger.DefaultLogger.Error(fmt.Errorf("配置秘钥认证失败！"))
			responses.ResponseApp.FailWithAgent(c, "", "配置秘钥认证失败！")
			c.Abort()
			return
		}
		c.Next()
	}
}
