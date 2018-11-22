package pipeline

import "testing"

type mock struct {}

func (m *mock) Execute() error {
	return nil
}

func TestPipelineCreation(t *testing.T) {
	pipeline := CreateJobPipe(10)

	if pipeline == nil {
		t.Error("pipeline creation failed")
	}
}

func TestSendingWorkToPipeline(t *testing.T) {
	pipeline := CreateJobPipe(10)
	jobs := make([]Job, 1)
	jobs[0] = &mock{}
	SendWorkToPipe(pipeline, &jobs)
}

func TestWorkerCreation(t *testing.T) {
	pipeline := CreateJobPipe(3)
	pipeline <- &mock{}
	pipeline <- &mock{}
	pipeline <- &mock{}
	close(pipeline)
	CreateWorkersForJobPipe(pipeline, 5)
}

