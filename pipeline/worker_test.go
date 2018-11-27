package pipeline

import "testing"

type mock struct {}

func (m *mock) Execute() error {
	return nil
}

func TestCanCreatePipelineConfig(t *testing.T){
	config := new(Config)
	if config == nil {
		t.Error("failed to create instance of pipeline config type")
	}
}

func TestPipelineCreation(t *testing.T) {
	pipeline := createJobPipe(10)

	if pipeline == nil {
		t.Error("pipeline creation failed")
	}
}

func TestSendingWorkToPipeline(t *testing.T) {
	pipeline := createJobPipe(10)
	jobs := make([]Job, 1)
	jobs[0] = &mock{}
	sendWorkToPipe(pipeline, &jobs)
}

func TestWorkerCreation(t *testing.T) {
	pipeline := createJobPipe(3)
	pipeline <- &mock{}
	pipeline <- &mock{}
	pipeline <- &mock{}
	close(pipeline)
	createWorkersForJobPipe(pipeline, 5)
}
