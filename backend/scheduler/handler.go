package scheduler

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"rlesjak.com/ha-scheduler/logs"
)

func handler(tag string, command string, job gocron.Job) {
	fmt.Println("------- command: " + command)
	fmt.Println("SCHEDULRE: " + tag + "\n" + job.Tags()[0])

	logs.Info.Printf("<JOB>{%s} Executed with command: (%s)", tag, command)
}
