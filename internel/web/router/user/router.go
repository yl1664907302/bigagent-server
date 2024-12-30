package user

import (
	"bigagent_server/internel/web/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (*UserRouter) Router(r gin.IRouter) {
	g := r.Group("/v1")
	UserApi := api.ApiGroupApp.UserApiGroup
	g.POST("/login", UserApi.Login)
	g.GET("/loginOut", UserApi.LoginOut)
}
