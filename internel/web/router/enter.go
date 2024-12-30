package router

import (
	"bigagent_server/internel/web/router/other"
	"bigagent_server/internel/web/router/server"
	"bigagent_server/internel/web/router/user"
)

type RouterGroup struct {
	ServerRouter server.ServerRouter
	OtherRouter  []other.OtherRouter
	UserRouter   user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
