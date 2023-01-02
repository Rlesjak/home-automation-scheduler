package services

import (
	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/services/parsers"
)

func CreateScheduledJob(condition string, command string, trigUuid uuid.UUID) (string, error) {

	parsedCondition, parseErr := parsers.ParseCondition(condition)
	if parseErr != nil {
		return "", parseErr
	}

	// Get scheduler
	schedule := scheduler.Scheduler
	// Set interval
	schedule.Every(parsedCondition.EveryNof)
	// Set event/unit
	buildEventsSchedule(parsedCondition.Events, schedule)

	// If at time of day is set, add it to schedule
	if parsedCondition.At.Valid {
		schedule.At(parsedCondition.At.String)
	}

	// Create job and return next time it will run
	job, err := scheduler.CreateJob(schedule, trigUuid, command)
	return job.NextRun().String(), err
}

func buildEventsSchedule(events []string, schedule *gocron.Scheduler) {
	for _, event := range events {
		switch event {
		// Setting units
		case "second":
			schedule.Second()
		case "minute":
			schedule.Minute()
		case "hour":
			schedule.Hour()
		case "day":
			schedule.Day()
		// Setting days of week
		case parsers.DaysOfWeek[0]:
			schedule.Day().Monday()
		case parsers.DaysOfWeek[1]:
			schedule.Day().Tuesday()
		case parsers.DaysOfWeek[2]:
			schedule.Day().Wednesday()
		case parsers.DaysOfWeek[3]:
			schedule.Day().Thursday()
		case parsers.DaysOfWeek[4]:
			schedule.Day().Friday()
		case parsers.DaysOfWeek[5]:
			schedule.Day().Saturday()
		case parsers.DaysOfWeek[6]:
			schedule.Day().Sunday()
		}
	}
}
