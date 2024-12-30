package server

import (
	"bigagent_server/internel/web/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerRouter struct{}

func (*ServerRouter) Router(r gin.IRouter) {
	g := r.Group("/v1")
	ServerApi := api.ApiGroupApp.ServerApiGroup
	g.GET("/agent_id", ServerApi.SearchAgent)
	g.GET("/agent_id_patrol", ServerApi.SearchAgentPatrol)
	g.POST("/push", ServerApi.PushAgentConfig)
	g.POST("/push_host", ServerApi.PushAgentConfigByHost)
	g.POST("/add", ServerApi.AddAgentConfig)
	g.GET("/get", ServerApi.GetAgentConfig)
	g.DELETE("/del", ServerApi.DelAgentConfig)
	g.PUT("/edit", ServerApi.EditAgentConfig)
	g.DELETE("/del_agent", ServerApi.DeleteAgentInfo)
	g.GET("/info", ServerApi.GetAgentInfo)
	g.GET("/info_sse", ServerApi.GetAgentInfoSSE)
	g.GET("/get_dead", ServerApi.GetAgentNumDead)
	g.GET("/get_cofail", ServerApi.GetAgentConfigFail)

	// swagger api docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
