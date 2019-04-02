package files

import (
	"io"

	ops "gitlab.com/davecremins/safe-deposit-box/io-ops"
	pipeline "gitlab.com/davecremins/safe-deposit-box/pipeline"
	workers "gitlab.com/davecremins/safe-deposit-box/workers"
)

type fileOp struct {
	handle io.ReaderAt
	data   *ops.Chunk
}

func (op *fileOp) Execute() error {
	ops.ReadIntoChunk(op.handle, op.data)
	return nil
}

// ReadInChunks calculates the offsets of the chunks, creates a pipeline configuration
// and executes it to have the workers read the chunks concurrently.
func ReadInChunks(reader io.ReaderAt, dataSize int64, bufferSize int) *[]ops.Chunk {
	chunks := ops.PrepareChunks(dataSize, bufferSize)
	chunkAmount := len(*chunks)
	pipelineConfig := pipeline.Config{
		JobSize:      chunkAmount,
		WorkerAmount: chunkAmount,
		Jobs:         allocateFileOps(reader, chunks),
		LoadBalance:  false,
	}
	pipeline.Create(pipelineConfig)
	return chunks
}

func allocateFileOps(reader io.ReaderAt, chunks *[]ops.Chunk) *[]workers.Job {
	amount := len(*chunks)
	fileOps := make([]workers.Job, amount)
	for i := 0; i < amount; i++ {
		fileOps[i] = &fileOp{handle: reader, data: &(*chunks)[i]}
	}
	return &fileOps
}
