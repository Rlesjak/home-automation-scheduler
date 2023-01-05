package scheduler

import (
	"fmt"

	"github.com/go-co-op/gocron"
)

type JobCommand struct {
	Run func()
}

// This method is called by the scheduler
// with the required parameters and at the required time
// can be overrided with RegisterJobHandler
var executeJob = func(tag string, command JobCommand, job gocron.Job) {
	fmt.Println("------- DEFAULT JOB HANDLER: ")
	fmt.Println("params: ")
	fmt.Println(tag)
	fmt.Println(command)
	fmt.Println(command)
}

func RegisterJobHandler(handler func(string, JobCommand, gocron.Job)) {
	executeJob = handler
}
