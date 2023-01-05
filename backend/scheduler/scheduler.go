package scheduler

import (
	"errors"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/gookit/config/v2"
)

//TODO: Create RegisterJobErrorLogger function

var scheduler *gocron.Scheduler

func InitScheduler(config config.Config) {
	scheduler = gocron.NewScheduler(time.Local)
	scheduler.StartAsync()
}

// Get scheduler instance
func GetScheduler() *gocron.Scheduler {
	return scheduler
}

// Get scheduler with already configured interval
// and register a job with the provided uuid
func CreateJob(scheduler *gocron.Scheduler, jobuuid uuid.UUID, jobcommand JobCommand) (*gocron.Job, error) {
	uuidString := jobuuid.String()

	// Check if job with the same ID already exists
	jobs, _ := scheduler.FindJobsByTag(uuidString)
	if len(jobs) > 0 {
		// If it does, return an error
		return nil, errors.New("Job '" + uuidString + "' already registered!")
	}

	// Start job with schedule
	job, err := scheduler.
		Tag(uuidString).
		WaitForSchedule().
		DoWithJobDetails(executeJob, uuidString, jobcommand)
	return job, err
}

func DeleteJob(jobuuid uuid.UUID) error {
	return scheduler.RemoveByTag(jobuuid.String())
}
