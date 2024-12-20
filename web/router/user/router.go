package user

import (
	"bigagent_server/web/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (*UserRouter) Router(r *gin.Engine) {
	g := r.Group("/v1")
	UserApi := api.ApiGroupApp.UserApiGroup
	g.POST("/login", UserApi.Login)
}
