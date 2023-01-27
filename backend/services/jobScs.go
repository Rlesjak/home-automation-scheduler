package services

import (
	"context"

	"github.com/google/uuid"
	db "rlesjak.com/ha-scheduler/database/generated"
	"rlesjak.com/ha-scheduler/logs"
	models "rlesjak.com/ha-scheduler/model"
)

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

// I dont think it is safe to run this function in a go routine
// since it uses global scheduler pointer
func CreateJobsFromAllActiveTriggers() error {
	triggers, err := models.Q.GetAllActiveTriggers(context.Background())
	if err != nil {
		logs.Error.Fatalln(err.Error())
		return err
	}

	// Create jobs from triggers list
	for _, trigger := range triggers {

		// Register job to the scheduler
		_, crTrigErr := CreateScheduledJob(
			trigger.Condition.String,
			trigger.Command.String,
			trigger.Uuid,
		)

		if crTrigErr != nil {
			logs.Error.Printf("JOB<%s> %s\n", trigger.Uuid.String(), crTrigErr.Error())
			// If trigger has an error creating a job, deactivate it
			models.Q.UpdateTriggerActive(
				context.Background(),
				db.UpdateTriggerActiveParams{
					Uuid:   trigger.Uuid,
					Active: false,
				},
			)

		} else {
			logs.Info.Printf("JOB<%s> Restored\n", trigger.Uuid.String())
		}
	}

	return nil
}
