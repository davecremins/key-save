package files

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type chunk struct {
	size   int
	offset int64
}

type job struct {
	handle io.ReaderAt
	data   *chunk
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

	chunks := prepareChunks(fileinfo.Size(), bufferSize)
	chunkAmount := len(*chunks)

	jobs := make(chan job, chunkAmount)
	go allocateJobs(file, chunks, chunkAmount, jobs)

	jobResult := make(chan *[]byte, chunkAmount)
	totalByteReadCount := make(chan int)
	go processResults(jobResult, totalByteReadCount)

	createWorkers(jobs, jobResult, chunkAmount)
	fmt.Println("--- Total amount of bytes read:", <-totalByteReadCount, " ---")
}

func allocateJobs(file io.ReaderAt, chunks *[]chunk, chunkAmount int, jobs chan<- job) {
	for i := 0; i < chunkAmount; i++ {
		jobs <- job{handle: file, data: &(*chunks)[i]}
	}
	close(jobs)
}

func processResults(jobResults <-chan *[]byte, done chan<- int) {
	totalByteCount := 0
	for bRead := range jobResults {
		totalByteCount += len(*bRead)
		fmt.Println("Bytes read:", string(*bRead))
	}
	done <- totalByteCount
}

func createWorkers(jobs chan job, jobResults chan *[]byte, chunkAmount int) {
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
		fmt.Println("Processing job in worker:", id)
		buffer := read(j.handle, *j.data)
		bytesRead <- &buffer
		fmt.Println("Finished processing job in worker:", id)
	}
	wg.Done()
}

func prepareChunks(blobSize int64, bufferSize int) *[]chunk {
	size := int(blobSize)
	parts := size / bufferSize
	chunks := make([]chunk, parts)

	for i := 0; i < parts; i++ {
		chunks[i].size = bufferSize
		chunks[i].offset = int64(bufferSize * i)
	}

	// Add the remaining number of bytes as last chunk size
	if remainder := size % bufferSize; remainder != 0 {
		c := chunk{size: remainder, offset: int64(parts * bufferSize)}
		chunks = append(chunks, c)
	}

	return &chunks
}

// TODO: Move this out into its own package
func read(handle io.ReaderAt, part chunk) []byte {
	buffer := make([]byte, part.size)
	_, err := handle.ReadAt(buffer, part.offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("fatal: should not have read EOF")
			panic(err)
		}
	}
	return buffer
}
