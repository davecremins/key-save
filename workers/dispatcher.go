package workers

import "sync"

type Dispatcher struct {
	WorkerPool   Pool
	WorkerAmount int
	Workers      []Worker
	WaitGroup    *sync.WaitGroup
}

func NewDispatcher(workerAmount int) *Dispatcher {
	pool := make(Pool, workerAmount)
	var wg sync.WaitGroup
	dispatcher := &Dispatcher{
		WorkerPool:   pool,
		WorkerAmount: workerAmount,
		WaitGroup:    &wg}

	return dispatcher
}

func (d *Dispatcher) CreateWorkers() {
	i := 0
	for i < d.WorkerAmount {
		i++
		d.WaitGroup.Add(1)
		worker := NewWorker(i, d.WorkerPool, d.WaitGroup)
		d.Workers = append(d.Workers, worker)
		worker.Start()
	}
}

func (d *Dispatcher) DispatchFrom(jobQueue JobChannel) {
	d.WaitGroup.Add(1)
	go func() {
		defer d.WaitGroup.Done()
		for job := range jobQueue {
			// Wait for available job channel and dispatch to it
			availableJobChannel := <-d.WorkerPool
			availableJobChannel <- job
		}

		for i := 0; i < d.WorkerAmount; i++ {
			jobChannel := <-d.WorkerPool
			close(jobChannel)
		}
	}()
}

func (d *Dispatcher) WaitForCompletion() {
	d.WaitGroup.Wait()
}
