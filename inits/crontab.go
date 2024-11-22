package inits

import "bigagent_server/utils/crontab"

func CronTask() {
	crontab.ScrapeCrontab()
}
