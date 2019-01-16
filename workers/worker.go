package workers

import (
	"fmt"
	"sync"
)

type Job interface {
	Execute() error
}

type JobChannel chan Job

type Pool chan JobChannel

type Worker struct {
	ID         int
	WorkerPool Pool
	JobChannel JobChannel
	WG         *sync.WaitGroup
}

func NewWorker(id int, pool Pool, wg *sync.WaitGroup) Worker {
	return Worker{
		ID:         id,
		WorkerPool: pool,
		JobChannel: make(JobChannel),
		WG:         wg}
}

func (w *Worker) Start() {
	go func() {
		workerStartMsg := fmt.Sprintf("Worker %d started", w.ID)
		fmt.Println(workerStartMsg)
		for {
			// Send the worker's JobChannel to the WorkerPool
			w.WorkerPool <- w.JobChannel
			job, ok := <-w.JobChannel
			if ok {
				// Job received from workers own channel
				fmt.Println("Processing job in worker", w.ID)
				job.Execute()
			} else {
				w.Stop()
				return
			}

		}
	}()
}

func (w *Worker) Stop() {
	fmt.Println("Shutting down worker", w.ID)
	w.WG.Done()
}
