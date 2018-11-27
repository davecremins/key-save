package pipeline

import (
	"fmt"
	"sync"
)

var showInfo = true

func info(args ...interface{}) {
	if showInfo {
		fmt.Println(args...)
	}
}

type Job interface {
	Execute() error
}

type Config struct {
	JobSize int
	WorkerAmount int
	Jobs *[]Job
	LoadBalance bool
}

func Create(config Config){
	jobCh := createJobPipe(config.JobSize)
	sendWorkToPipe(jobCh, config.Jobs)
	createWorkersForJobPipe(jobCh, config.WorkerAmount)
}

func createJobPipe(size int) chan Job {
	return make(chan Job, size)
}

func createWorkersForJobPipe(jobPipe chan Job, workerCount int) {
	fmt.Sprintf("Creating %d workers for pipeline", workerCount)
	var wg sync.WaitGroup
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go createWorker(w+1, jobPipe, &wg)
	}
	wg.Wait()
}

func sendWorkToPipe(jobPipe chan<- Job, jobs *[]Job){
	go func () {
		for _, work := range (*jobs) {
			jobPipe <- work
		}
		close(jobPipe)
	}()
}

func createWorker(id int, jobPipe <-chan Job, wg *sync.WaitGroup) {
	for work := range jobPipe {
		info("Processing job in worker:", id)
		work.Execute()
		info("Finished processing job in worker:", id)
	}
	wg.Done()
}
