package services

import (
	"github.com/gookit/config/v2"
	"rlesjak.com/ha-scheduler/scheduler"
)

func InitialiseServices(config config.Config) error {
	scheduler.RegisterJobHandler(executeTrigger)

	// Create Jobs from all active triggers on app startup
	// ( Restore them after server restart )
	if restoreFatalErr := CreateJobsFromAllActiveTriggers(); restoreFatalErr != nil {
		return restoreFatalErr
	}

	return nil
}
