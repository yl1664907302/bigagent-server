package services

import (
	"bigagent_server/internel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentService interface {
	GetAgentInfo(c *gin.Context) ([]model.AgentInfo, error)
	AddAgentConfig(c *gin.Context) error
	EditAgentConfig(c *gin.Context) error
	DelAgentConfig(c *gin.Context) error
	GetAgentNum(c *gin.Context) (int, error)
	GetAgentNumDead2Live(c *gin.Context) (int, int, error)
	SearchAgentNet(c *gin.Context) (string, error)
	GetAgentConfigs2num(c *gin.Context) ([]model.AgentConfigDB, int, error)
	GetAgentConfig2Nets(c *gin.Context, body []byte) (*model.AgentConfigDB, []string, string, bool, error)
	GetAgentConfig2Uuids(c *gin.Context, body []byte, key bool) (*model.AgentConfigDB, []string, string, error)
	GetAgentRedictShow(c *gin.Context, host string, key string, action bool) (*http.Response, error)
	UpdateAgentConfigStatus(c *gin.Context, id int, status string) error
	DeleteAgentInfo(c *gin.Context) error
	GetAgentConfigNEW2Fail(c *gin.Context) (int, int, error)
}
