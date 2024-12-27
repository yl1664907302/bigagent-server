package other

import (
	"bigagent_server/strategy"
	"bigagent_server/web/api"
	"github.com/gin-gonic/gin"
)

type OtherRouter struct {
	K bool
	S *strategy.CmdbServe
}

func (*OtherRouter) Router(path string, r gin.IRouter) {
	g := r.Group("/" + path)
	OtherApi := api.ApiGroupApp.OtherApiGroup
	g.GET("/showdata", OtherApi.Showdata)
}
