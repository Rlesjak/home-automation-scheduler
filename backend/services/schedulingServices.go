package services

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/services/parsers"
)

func CreateScheduledJob(condition string, command string, trigUuid uuid.UUID) (string, error) {

	// Parse given condition
	parsedCondition, conditionParseErr := parsers.ParseCondition(condition)
	if conditionParseErr != nil {
		return "", conditionParseErr
	}

	// Parse given command string
	parsedCommand, commandParseErr := parsers.ParseCommand(command)
	if commandParseErr != nil {
		return "", commandParseErr
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
	return job.NextRun().String(), err
}

// This method is called by the scheduler
// with the required parameters and at the required time
func executeTrigger(triggerUuid string, command scheduler.JobCommand, job gocron.Job) {
	fmt.Println("EXECTUE TRIGGER")
	command.Run()
	// fmt.Println("SCHEDULRE: " + triggerUuid + "\n" + job.Tags()[0])

	// logs.Info.Printf("<JOB>{%s} Executed with command: (%s)", triggerUuid, command)
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
