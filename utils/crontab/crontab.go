package crontab

import (
	"bigagent_server/db/mysqldb"
	"bigagent_server/utils/logger"
	"github.com/robfig/cron/v3"
	"time"
)

// CronTask Crontab执行的任务列表
func CronTask() {
	checkAgentStatus()
}

// ScrapeCrontab 初始化采集crontab任务
func ScrapeCrontab() {
	crontabRule := "@every 5s"
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
	// 超时时间阈值，例如 15 秒未通信
	timeout := now.Add(-10 * time.Second)
	// 更新掉线的 Agent
	_, err := mysqldb.AgentUpdateActiveToDead(timeout)
	if err != nil {
		logger.DefaultLogger.Error("sql执行失败：", err)
	}
	//for _, dagent := range dagents {
	//	logger.DefaultLogger.Infof("agent端：%s,已掉线", dagent.IPv4First)
	//}
}
