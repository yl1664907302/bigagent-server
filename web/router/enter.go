package router

import (
	"bigagent_server/web/router/other"
	"bigagent_server/web/router/server"
)

type RouterGroup struct {
	ServerRouter server.ServerRouter
	OtherRouter  []other.OtherRouter
}

var RouterGroupApp = new(RouterGroup)
