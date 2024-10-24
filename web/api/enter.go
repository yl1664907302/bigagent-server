package api

import "bigagent_server/web/api/server"

type ApiGroup struct {
	ServerApiGroup server.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
