package services

import (
	"bigagent_server/model"
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
	GetAgentConfig2Nets(c *gin.Context) (*model.AgentConfigDB, []string, error)
	GetAgentConfig2Uuids(c *gin.Context) (*model.AgentConfigDB, []string, error)
	GetAgentRedict(c *gin.Context, host string) (*http.Response, error)
	UpdateAgentConfigTimes(c *gin.Context, id int) error
}
