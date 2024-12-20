package server

import (
	"bigagent_server/web/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerRouter struct{}

func (*ServerRouter) Router(r *gin.Engine) {
	g := r.Group("/v1")
	ServerApi := api.ApiGroupApp.ServerApiGroup
	g.GET("/agent_id", ServerApi.SearchAgent)
	g.POST("/push", ServerApi.PushAgentConfig)
	g.POST("/push_host", ServerApi.PushAgentConfigByHost)
	g.POST("/add", ServerApi.AddAgentConfig)
	g.GET("/get", ServerApi.GetAgentConfig)
	g.DELETE("/del", ServerApi.DelAgentConfig)
	g.PUT("/edit", ServerApi.EditAgentConfig)
	g.GET("/info", ServerApi.GetAgentInfo)
	g.GET("/info_sse", ServerApi.GetAgentInfoSSE)

	// swagger api docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
