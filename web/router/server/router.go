package server

import (
	"bigagent_server/web/api"
	"github.com/gin-gonic/gin"
)

type ServerRouter struct{}

func (*ServerRouter) Router(r *gin.Engine) {
	g := r.Group("/v1")
	ServerApi := api.ApiGroupApp.ServerApiGroup
	g.GET("/agent_id", ServerApi.SearchAgent)
	g.POST("/push", ServerApi.PushAgentData)
}
