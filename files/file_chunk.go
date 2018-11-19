package files

import (
	"fmt"
	"io"
	"os"
	"sync"

	ops "gitlab.com/davecremins/safe-deposit-box/io-ops"
)

type job struct {
	handle io.ReaderAt
	data   *ops.Chunk
}

var showInfo = true

func info(args ...interface{}) {
	if showInfo {
		fmt.Println(args...)
	}
}

func ReadFileInChunks(filepath string, bufferSize int) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	readInChunks(file, fileinfo.Size(), bufferSize)
}

func readInChunks(file io.ReaderAt, dataSize int64, bufferSize int) *[]ops.Chunk {
	chunks := ops.PrepareChunks(dataSize, bufferSize)
	chunkAmount := len(*chunks)
	jobs := make(chan job, chunkAmount)
	go allocateJobs(file, chunks, chunkAmount, jobs)

	createWorkers(jobs, chunkAmount)
	return chunks
}

func allocateJobs(file io.ReaderAt, chunks *[]ops.Chunk, chunkAmount int, jobs chan<- job) {
	for i := 0; i < chunkAmount; i++ {
		jobs <- job{handle: file, data: &(*chunks)[i]}
	}
	close(jobs)
}

func processResults(jobResults <-chan *[]byte, done chan<- int) {
	totalByteCount := 0
	for bRead := range jobResults {
		totalByteCount += len(*bRead)
		info("Bytes read:", string(*bRead))
	}
	done <- totalByteCount
}

func createWorkers(jobs chan job, chunkAmount int) {
	info(fmt.Sprintf("Creating %d workers to read file", chunkAmount))
	var wg sync.WaitGroup
	for w := 0; w < chunkAmount; w++ {
		wg.Add(1)
		go readWorker(w+1, jobs, &wg)
	}
	wg.Wait()
}

func readWorker(id int, jobs <-chan job, wg *sync.WaitGroup) {
	for j := range jobs {
		info("Processing job in worker:", id)
		ops.ReadIntoChunk(j.handle, j.data)
		info("Finished processing job in worker:", id)
	}
	wg.Done()
}
