package workers

import "testing"

func TestCreationOfDispatcher(t *testing.T) {
	workerCount := 5
	dispatcher := NewDispatcher(workerCount)
	if cap(dispatcher.WorkerPool) != workerCount {
		t.Error("worker pool capacity is incorrect")
	}
}

func TestCreationOfWorkersFromDispatcher(_ *testing.T) {
	dispatcher := NewDispatcher(5)
	dispatcher.CreateWorkers()
}

func TestDispatchFromDispatcher(_ *testing.T) {
	dispatcher := NewDispatcher(5)
	jobQueue := make(chan Job, 5)
	dispatcher.DispatchFrom(jobQueue)
}
