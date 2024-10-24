package router

import "bigagent_server/web/router/server"

type RouterGroup struct {
	ServerRouter server.ServerRouter
}

var RouterGroupApp = new(RouterGroup)
