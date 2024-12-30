package api

import (
	"bigagent_server/internel/web/api/other"
	"bigagent_server/internel/web/api/server"
	"bigagent_server/internel/web/api/user"
)

type ApiGroup struct {
	ServerApiGroup server.ApiGroup
	OtherApiGroup  other.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
