package utils

import (
	"github.com/robfig/cron/v3"
)

func NewCrond(stime string, send func()) {

	crontab := cron.New(cron.WithSeconds())
	defer crontab.Stop()

	crontab.AddFunc(stime, send)

	crontab.Start()

	select {}
}
