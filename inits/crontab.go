package inits

import (
	"bigagent_server/internel/utils/crontab"
)

func CronTask() {
	crontab.ScrapeCrontab()
}
