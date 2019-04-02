package workers

import (
	"fmt"
	"sync"
)

// Job is an interface abstraction for anything that implements the Execute function.
type Job interface {
	Execute() error
}

// JobChannel is a channel that accepts Job types.
type JobChannel chan Job

// Pool is a channel that accepts JobChannel types.
type Pool chan JobChannel

// Worker is a structure that executes Job types.
type Worker struct {
	ID         int
	WorkerPool Pool
	JobChannel JobChannel
	WG         *sync.WaitGroup
}

// NewWorker creates a new Worker type.
func NewWorker(id int, pool Pool, wg *sync.WaitGroup) Worker {
	return Worker{
		ID:         id,
		WorkerPool: pool,
		JobChannel: make(JobChannel),
		WG:         wg}
}

// Start is a gorountine that sends the worker's job channel to the pool
// in order for the consumer to dispatch a job to the worker. Once the consumer
// closes the worker's job channel, the gorountine exits.
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

// Stop executes the Done() function on the worker's wait group
// to signal it is finished.
func (w *Worker) Stop() {
	fmt.Println("Shutting down worker", w.ID)
	w.WG.Done()
}
