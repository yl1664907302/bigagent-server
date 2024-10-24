package inits

import (
	"bigagent_server/web/router"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r1 := router.RouterGroupApp.ServerRouter
	r1.Router(r)
	return r
}
