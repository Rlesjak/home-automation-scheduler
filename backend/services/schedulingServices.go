package services

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"rlesjak.com/ha-scheduler/logs"
	models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/services/parsers"
)

func CreateScheduledJob(condition string, command string, trigUuid uuid.UUID) (string, error) {

	parsedCondition, parsedCommand, parserError := parsers.ParseJob(condition, command)
	if parserError != nil {
		return "", parserError
	}

	//*************** IMPORTANT!!!
	// DO NOT RETURN THIS FUNCTION ANYWHERE FROM HERE
	// TO THE scheduler.CreateJob() CALL
	// Because global instance of scheduler is used
	// if the process of building a schedule gets interrupted
	// there will be mixing of schedules from multiple calls

	// Get scheduler
	schedule := scheduler.GetScheduler()
	// Set interval
	schedule.Every(parsedCondition.EveryNof)
	// Set event/unit
	buildEventsSchedule(parsedCondition.Events, schedule)
	// If at time of day is set, add it to schedule
	if parsedCondition.At.Valid {
		schedule.At(parsedCondition.At.String)
	}
	// Create job and return next time it will run
	job, err := scheduler.CreateJob(schedule, trigUuid, parsedCommand)

	if err != nil {
		return "", err
	}

	return job.NextRun().String(), err
}

func CreateScheduledJobFromTriggerUuid(ctx context.Context, trigUuid uuid.UUID) (string, error) {
	// Fetch trigger data from database
	trigger, qerr := models.Q.GetTriggerByUuid(ctx, trigUuid)
	if qerr != nil {
		return "", qerr
	}

	//TODO: Implement logic if trigger is runnign an element script
	// If elementuuid is not present, all needed info is already available
	// if trigger.ElementUuid.Valid == false {
	return CreateScheduledJob(
		trigger.Condition.String,
		trigger.Command.String,
		trigger.Uuid,
	)
	// } else {
	// 	// Fetch element info
	// 	element, err := models.Q.GetElementById....
	// }
}

// This method is called by the scheduler
// with the required parameters and at the required time
func executeTrigger(triggerUuid string, command scheduler.JobCommand, job gocron.Job) {
	// Run the command
	command.Run()

	// Log
	logs.Info.Printf("<JOB>{%s} Executed", triggerUuid)
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
