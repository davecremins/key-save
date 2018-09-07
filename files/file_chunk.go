package files

import (
	"fmt"
	"io"
	"os"
)

type chunk struct {
	size   int
	offset int64
}

type readJob struct {
	handle io.ReaderAt
	data   *chunk
}

func ChunkFile(filepath string, bufferSize int) {
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

	jobs := make(chan readJob)
	jobResult := make(chan *[]byte)
	for w := 0; w < chunkAmount; w++ {
		go readWorker(w+1, jobs, jobResult)
	}

	for i := 0; i < chunkAmount; i++ {
		jobs <- readJob{handle: file, data: &(*chunks)[i]}
	}
	close(jobs)

	for bRead := range jobResult {
		fmt.Println("Bytes read:", string(*bRead))
	}
}

func readWorker(id int, jobs <-chan readJob, bytesRead chan<- *[]byte) {
	for j := range jobs {
		fmt.Println("Processing job in worker:", id)
		buffer := read(j.handle, *j.data)
		bytesRead <- &buffer
	}
}

func prepareChunks(blobSize int64, bufferSize int) *[]chunk {
	size := int(blobSize)
	parts := size / bufferSize
	chunks := make([]chunk, parts)

	for i := 0; i < parts; i++ {
		chunks[i].size = bufferSize
		chunks[i].offset = int64(bufferSize * i)
	}

	// Add the remaining  number of bytes as last chunk size
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
