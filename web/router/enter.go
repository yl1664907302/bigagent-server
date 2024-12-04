package router

import (
	"bigagent_server/web/router/other"
	"bigagent_server/web/router/server"
	"bigagent_server/web/router/user"
)

type RouterGroup struct {
	ServerRouter server.ServerRouter
	OtherRouter  []other.OtherRouter
	UserRouter   user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
