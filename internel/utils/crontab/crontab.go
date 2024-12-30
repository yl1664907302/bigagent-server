package crontab

import (
	"bigagent_server/internel/config"
	"bigagent_server/internel/db/mysqldb"
	"bigagent_server/internel/logger"
	"github.com/robfig/cron/v3"
	"time"
)

// CronTask Crontab执行的任务列表
func CronTask() {
	checkAgentStatus()
}

// ScrapeCrontab crontab任务
func ScrapeCrontab() {
	crontabRule := "@every " + config.CONF.System.Times
	c := cron.New()

	addFunc, err := c.AddFunc(crontabRule, CronTask)
	if err != nil {
		logger.DefaultLogger.Error("定时任务启动异常：", err)
		return
	}
	c.Start()
	logger.DefaultLogger.Info("定时任务启动成功,EntryID：", addFunc)
}

func checkAgentStatus() {
	// 当前时间
	now := time.Now()
	// 超时时间阈值，例如 10 秒未通信
	out := config.CONF.System.Agent_outtime
	timeout := now.Add(-time.Duration(out) * time.Second)
	// 更新掉线的 Agent
	_, err := mysqldb.AgentUpdateActiveToDead(timeout)
	if err != nil {
		logger.DefaultLogger.Error("sql执行失败：", err)
	}
	//for _, dagent := range dagents {
	//	logger.DefaultLogger.Infof("agent端：%s,已掉线", dagent.IPv4First)
	//}
}
