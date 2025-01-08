package services

import (
	conf "bigagent_server/internel/config"
	dao "bigagent_server/internel/db/mysqldb"
	"bigagent_server/internel/db/redis"
	"bigagent_server/internel/logger"
	"bigagent_server/internel/model"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
)

type AgentServiceImpV1 struct{}

func (s *AgentServiceImpV1) GetAgentConfigNEW2Fail(c *gin.Context) (int, int, error) {
	id, err := dao.AgentconfigNewID()
	if err != nil {
		logger.DefaultLogger.Error(err)
		return 0, 0, err
	}
	fnum, err := dao.AgentconfigSelectByFail(id)
	if err != nil {
		return 0, 0, err
	}
	i, err := strconv.Atoi(id)
	return i, fnum, err
}

func (s *AgentServiceImpV1) DeleteAgentInfo(c *gin.Context) error {
	id := c.Query("uuid")
	err := dao.AgentDelete(id)
	if err != nil {
		logger.DefaultLogger.Errorf(err.Error())
		return err
	}
	return nil
}

func (s *AgentServiceImpV1) GetAgentNumDead2Live(c *gin.Context) (int, int, error) {
	dnum, anum, err := dao.AgentSelectlive2dead()
	if err != nil {
		return 0, 0, err
	}
	return dnum, anum, err
}

func (s *AgentServiceImpV1) GetAgentConfig2Nets(c *gin.Context) (*model.AgentConfigDB, []string, error) {
	// 获取配置ID并查询配置
	var requestdata map[string]int
	if body, err := c.GetRawData(); err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	} else if err = json.Unmarshal(body, &requestdata); err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	}

	id := requestdata["config_id"]
	config, err := dao.AgentconfigSelect(id)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	}

	// 获取agent地址列表
	agentAddrs, err := redisdb.ScanAgentAddresses(c)
	if err != nil || len(agentAddrs) == 0 {
		if err = dao.UpdateAgentAddressesToRedis(c); err != nil {
			logger.DefaultLogger.Error(err)
			return nil, nil, err
		}
		agentAddrs, _ = redisdb.ScanAgentAddresses(c)
	}
	return &config, agentAddrs, nil
}

func (s *AgentServiceImpV1) UpdateAgentConfigTimes(c *gin.Context, id int) error {
	err := dao.AgentconfigUpdateTimes(id)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	return nil
}

func (s *AgentServiceImpV1) GetAgentConfig2Uuids(c *gin.Context) (*model.AgentConfigDB, []string, error) {
	var requestData struct {
		ConfigID int      `json:"config_id"`
		Uuids    []string `json:"uuids"` // 主机IP列表
	}
	if body, err := c.GetRawData(); err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	} else if err = json.Unmarshal(body, &requestData); err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	}
	config, err := dao.AgentconfigSelect(requestData.ConfigID)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return nil, nil, err
	}
	return &config, requestData.Uuids, nil
}

func (s *AgentServiceImpV1) GetAgentConfigs2num(c *gin.Context) ([]model.AgentConfigDB, int, error) {
	configs, err := dao.AgentconfigSelectAll(c.Query("page"), c.Query("pageSize"))
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	num, err := dao.AgentconfigNetNum()
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	return configs, num, nil
}

func (s *AgentServiceImpV1) GetAgentRedict(c *gin.Context, host string, key string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(c.Request.Method, "http://"+host+"/"+c.Query("model_name")+"/"+key, c.Request.Body)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return nil, err
	}
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Add("Authorization", conf.CONF.System.Serct)
	// 此处密钥配置会在前端编写完后替换，改为从server的uri获取
	//req.Header.Add("Authorization", c.Request.Header.Get("Authorization"))
	resp, err := client.Do(req)
	return resp, nil
}

func (s *AgentServiceImpV1) SearchAgentNet(c *gin.Context) (string, error) {
	ip, err := dao.AgentNetIPSelectByUuid(c.Query("uuid"))
	if err != nil {
		logger.DefaultLogger.Error(err)
		return "", err
	}
	return ip, nil
}

func (s *AgentServiceImpV1) GetAgentInfo(c *gin.Context) ([]model.AgentInfo, error) {
	var agentInfos []model.AgentInfo
	if c.Query("type") == "" && c.Query("platform") == "" && c.Query("ip") == "" && c.Query("uuid") == "" && c.Query("active") == "" && c.Query("c_desc") == "" && c.Query("c_desc_f") == "" {
		s, err := dao.AgentInfoSelectAll(c.Query("page"), c.Query("pageSize"))
		agentInfos = s
		if err != nil {
			logger.DefaultLogger.Error(err)
			return nil, err
		}
	} else {
		s, err := dao.AgentInfoSelectByKeys(c.Query("page"), c.Query("pageSize"), c.Query("uuid"), c.Query("ip"), c.Query("type"), c.Query("platform"), c.Query("active"), c.Query("c_desc"), c.Query("c_desc_f"))
		agentInfos = s
		if err != nil {
			logger.DefaultLogger.Error(err)
			return nil, err
		}
	}
	return agentInfos, nil
}

func (s *AgentServiceImpV1) GetAgentNum(c *gin.Context) (int, error) {
	return dao.AgentNum()
}

func (s *AgentServiceImpV1) AddAgentConfig(c *gin.Context) error {
	body, err := c.GetRawData()
	var configDB model.AgentConfigDB
	err = json.Unmarshal(body, &configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	configDB.Status = "有效"
	err = dao.AgentconfigCreate(configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	return nil
}

func (s *AgentServiceImpV1) EditAgentConfig(c *gin.Context) error {
	body, err := c.GetRawData()
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	var configDB model.AgentConfigDB
	err = json.Unmarshal(body, &configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
	}
	err = dao.AgentconfigEdit(configDB.ID, configDB)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	return nil
}

func (s *AgentServiceImpV1) DelAgentConfig(c *gin.Context) error {
	id := c.Param("config_id")
	err := dao.AgentconfigDel(id)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	return nil
}

func (s *AgentServiceImpV1) UpdateAgentConfigStatus(c *gin.Context, id int, status string) error {
	err := dao.AgentconfigStatusChange(id, status)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return err
	}
	return nil
}

var AgentServiceImpV1App = new(AgentServiceImpV1)
