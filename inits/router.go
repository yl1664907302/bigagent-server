package inits

import (
	"bigagent_server/config"
	"bigagent_server/logger"
	responses "bigagent_server/web/response"
	"bigagent_server/web/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const shutdownTimeoutSecond = 10

func Router(r gin.IRouter) {
	r.Use(logger.GinLoggerMiddleware())
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
}

func Run(c context.Context, router func(gin.IRouter)) {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()
	g.Use(AuthMiddleware(), InitCORS(), gin.Recovery())
	router(&g.RouterGroup)

	// 创建 http.Server
	srv := &http.Server{
		Addr:    config.CONF.System.Addr, // 根据需要设置地址和端口
		Handler: g,
	}

	// 启动服务器
	go func() {
		logger.DefaultLogger.Info(fmt.Sprintf("服务器启动成功，监听地址: %s", config.CONF.System.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.DefaultLogger.Error(fmt.Errorf("服务器启动失败: %v", err))
		}
	}()

	// 阻塞主 goroutine，直到接收到退出信号
	<-c.Done()
	timeoutContext, cancel := context.WithTimeout(context.Background(), shutdownTimeoutSecond*time.Second)
	defer cancel()
	if err := srv.Shutdown(timeoutContext); err != nil {
		logger.DefaultLogger.Error("Server Shutdown", err)
		os.Exit(5)
	}
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
