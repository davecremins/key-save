package files

import (
	"fmt"
	"io"
	"os"

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

func ReadFileInChunks(filepath string, bufferSize int) *[]ops.Chunk {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return readInChunks(file, fileinfo.Size(), bufferSize)
}

func readInChunks(file io.ReaderAt, dataSize int64, bufferSize int) *[]ops.Chunk {
	chunks := ops.PrepareChunks(dataSize, bufferSize)
	chunkAmount := len(*chunks)
	pipelineConfig := pipeline.Config{
		JobSize:      chunkAmount,
		WorkerAmount: 10,
		Jobs:         allocateFileOps(file, chunks),
		LoadBalance:  false,
	}
	pipeline.Create(pipelineConfig)
	return chunks
}

func allocateFileOps(file io.ReaderAt, chunks *[]ops.Chunk) *[]pipeline.Job {
	amount := len(*chunks)
	fileOps := make([]pipeline.Job, amount)
	for i := 0; i < amount; i++ {
		fileOps[i] = &fileOp{handle: file, data: &(*chunks)[i]}
	}
	return &fileOps
}

func processResults(jobResults <-chan *[]byte, done chan<- int) {
	totalByteCount := 0
	for bRead := range jobResults {
		totalByteCount += len(*bRead)
		fmt.Println("Bytes read:", string(*bRead))
	}
	done <- totalByteCount
}
