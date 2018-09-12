package files

import (
	"fmt"
	"io"
	"os"
	"sync"

	ops "github.com/davecremins/safe-deposit-box/io-ops"
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

	chunks := ops.PrepareChunks(fileinfo.Size(), bufferSize)
	chunkAmount := len(*chunks)

	jobs := make(chan job, chunkAmount)
	go allocateJobs(file, chunks, chunkAmount, jobs)

	jobResult := make(chan *[]byte, chunkAmount)
	totalByteReadCount := make(chan int)
	go processResults(jobResult, totalByteReadCount)

	createWorkers(jobs, jobResult, chunkAmount)
	totalRead := <-totalByteReadCount
	info("--- Total amount of bytes read:", totalRead, " ---")
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

func createWorkers(jobs chan job, jobResults chan *[]byte, chunkAmount int) {
	info(fmt.Sprintf("Creating %d workers to read file", chunkAmount))
	var wg sync.WaitGroup
	for w := 0; w < chunkAmount; w++ {
		wg.Add(1)
		go readWorker(w+1, jobs, jobResults, &wg)
	}
	wg.Wait()
	close(jobResults)
}

func readWorker(id int, jobs <-chan job, bytesRead chan<- *[]byte, wg *sync.WaitGroup) {
	for j := range jobs {
		info("Processing job in worker:", id)
		buffer := ops.Read(j.handle, *j.data)
		bytesRead <- buffer
		info("Finished processing job in worker:", id)
	}
	wg.Done()
}
