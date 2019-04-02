// Package workers contains functions that allow for work to be dispatched for processing
package workers

import "sync"

// A Dispatcher manages the distribution of work to workers
type Dispatcher struct {
	WorkerPool   Pool
	WorkerAmount int
	Workers      []Worker
	WaitGroup    *sync.WaitGroup
}

// NewDispatcher creates a dispatcher and captures the amount of workers requested
func NewDispatcher(workerAmount int) *Dispatcher {
	pool := make(Pool, workerAmount)
	var wg sync.WaitGroup
	dispatcher := &Dispatcher{
		WorkerPool:   pool,
		WorkerAmount: workerAmount,
		WaitGroup:    &wg}

	return dispatcher
}

// CreateWorkers creates the requested amount of workers and adds
// them to an internal cache. The internal waitgroup is incremented
// for each worker created in order to wait for completion of all workers.
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

// DispatchFrom dispatches work from the job queue to the workers.
// Each worker has their own work channel and these channels are
// fed to the pool looking for work. Once a channel is retreived from
// the pool, work is submitted to the worker that owns that channel.
//
// Once all work from the job queue is complete, all channels from
// the workers are closed.
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

// WaitForCompletion waits for all work to be completed
func (d *Dispatcher) WaitForCompletion() {
	d.WaitGroup.Wait()
}
