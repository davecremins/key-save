package files

import (
	"io"

	ops "gitlab.com/davecremins/safe-deposit-box/io-ops"
	pipeline "gitlab.com/davecremins/safe-deposit-box/pipeline"
)

type fileOp struct {
	handle io.ReaderAt
	data   *ops.Chunk
}

func (op *fileOp) Execute() error {
	ops.ReadIntoChunk(op.handle, op.data)
	return nil
}

func ReadInChunks(reader io.ReaderAt, dataSize int64, bufferSize int) *[]ops.Chunk {
	chunks := ops.PrepareChunks(dataSize, bufferSize)
	chunkAmount := len(*chunks)
	pipelineConfig := pipeline.Config{
		JobSize:      chunkAmount,
		WorkerAmount: 10,
		Jobs:         allocateFileOps(reader, chunks),
		LoadBalance:  false,
	}
	pipeline.Create(pipelineConfig)
	return chunks
}

func allocateFileOps(reader io.ReaderAt, chunks *[]ops.Chunk) *[]pipeline.Job {
	amount := len(*chunks)
	fileOps := make([]pipeline.Job, amount)
	for i := 0; i < amount; i++ {
		fileOps[i] = &fileOp{handle: reader, data: &(*chunks)[i]}
	}
	return &fileOps
}
