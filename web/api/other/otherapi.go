package other

import (
	"bigagent_server/strategy"
	"github.com/gin-gonic/gin"
)

type OtherApi struct {
	K string
	S *strategy.CmdbServe
}

func (a *OtherApi) Showdata(c *gin.Context) {

}
