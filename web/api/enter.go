package api

import (
	"bigagent_server/web/api/other"
	"bigagent_server/web/api/server"
)

type ApiGroup struct {
	ServerApiGroup server.ApiGroup
	OtherApiGroup  other.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
