package scheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/gookit/config/v2"
)

//TODO: Create RegisterJobErrorLogger function

var Scheduler *gocron.Scheduler

func InitScheduler(config config.Config) {
	Scheduler = gocron.NewScheduler(time.Local)
	Scheduler.StartAsync()
}

func handler(tag string, command string, job gocron.Job) {
	fmt.Println("-------")
	fmt.Println("SCHEDULRE: " + tag + "\n" + job.Tags()[0])
}

// Get scheduler with already configured interval
// and register a job with the provided uuid
func CreateJob(scheduler *gocron.Scheduler, jobuuid uuid.UUID, jobcommand string) (*gocron.Job, error) {
	uuidString := jobuuid.String()

	// Check if job with the same ID already exists
	jobs, _ := scheduler.FindJobsByTag(uuidString)
	if len(jobs) > 0 {
		// If it does, return an error
		return nil, errors.New("Job '" + uuidString + "' already registered!")
	}

	job, err := scheduler.
		Tag(uuidString).
		StartImmediately().
		DoWithJobDetails(handler, uuidString, jobcommand)
	return job, err
}

func DeleteJob(jobuuid uuid.UUID) error {
	return Scheduler.RemoveByTag(jobuuid.String())
}
