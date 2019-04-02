package pipeline

import (
	"fmt"

	workers "gitlab.com/davecremins/safe-deposit-box/workers"
)

var showInfo = true

func info(args ...interface{}) {
	if showInfo {
		fmt.Println(args...)
	}
}

// Config contains the configuration for a pipeline.
type Config struct {
	JobSize      int
	WorkerAmount int
	Jobs         *[]workers.Job
	LoadBalance  bool
}

// Create initialises a new pipeline by creating a job channel,
// sending work to it and creating workers to process work on the channel.
func Create(config Config) {
	jobCh := createJobPipe(config.JobSize)
	sendWorkToPipe(jobCh, config.Jobs)
	createWorkersForJobPipe(jobCh, config.WorkerAmount)
}

func createJobPipe(size int) workers.JobChannel {
	return make(workers.JobChannel, size)
}

func createWorkersForJobPipe(jobPipe workers.JobChannel, workerCount int) {
	dispatcher := workers.NewDispatcher(workerCount)
	dispatcher.CreateWorkers()
	dispatcher.DispatchFrom(jobPipe)
	dispatcher.WaitForCompletion()
	info("Pipeline complete")
}

func sendWorkToPipe(jobPipe chan<- workers.Job, jobs *[]workers.Job) {
	go func() {
		for _, work := range *jobs {
			jobPipe <- work
		}
		close(jobPipe)
	}()
}
