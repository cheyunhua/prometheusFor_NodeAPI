package utils

import (
	"awesomeProject/logger"
	"github.com/robfig/cron/v3"
)

func NewCrond(stime string, send func()) {

	crontab := cron.New(cron.WithSeconds())
	defer crontab.Stop()

	_, err := crontab.AddFunc(stime, send)
	if err != nil {
		logger.DefaultLogger.Errorf("%+v", err)
	}

	crontab.Start()

	select {}
}
