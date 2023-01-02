package services

import (
	"github.com/google/uuid"
	"rlesjak.com/ha-scheduler/scheduler"
)

func CreateScheduledJob(condition string, command string, trigUuid uuid.UUID) error {
	schedule := scheduler.Scheduler.Every(2).Seconds()
	_, err := scheduler.CreateJob(schedule, trigUuid, command)
	return err
}
