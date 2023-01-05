package services

import (
	"github.com/gookit/config/v2"
	"rlesjak.com/ha-scheduler/scheduler"
)

func InitialiseServices(config config.Config) {
	scheduler.RegisterJobHandler(executeTrigger)
}
