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

	filesize := int(fileinfo.Size())
	fileParts := filesize / bufferSize
	chunks := make([]chunk, fileParts)

	for i := 0; i < fileParts; i++ {
		chunks[i].size = bufferSize
		chunks[i].offset = int64(bufferSize * i)
	}

	// Add the remainder number of bytes as last chunk size
	if remainder := filesize % bufferSize; remainder != 0 {
		c := chunk{size: remainder, offset: int64(fileParts * bufferSize)}
		fileParts++
		chunks = append(chunks, c)
	}

	var wg sync.WaitGroup
	wg.Add(fileParts)

	for i := 0; i < fileParts; i++ {
		go read(file, chunks[i], wg)
	}

	wg.Wait()
}

// TODO: Move this out into its own package and remove dependancy on WaitGroup
func read(handle io.ReaderAt, part chunk, wg sync.WaitGroup) []byte {
	defer wg.Done()

	buffer := make([]byte, part.size)
	_, err := handle.ReadAt(buffer, part.offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("fatal: should not have read EOF")
			panic(err)
		}
	}
	// Use a channel here to send buffer
	return buffer
}
