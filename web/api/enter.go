package api

import (
	"bigagent_server/web/api/other"
	"bigagent_server/web/api/server"
	"bigagent_server/web/api/user"
)

type ApiGroup struct {
	ServerApiGroup server.ApiGroup
	OtherApiGroup  other.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
