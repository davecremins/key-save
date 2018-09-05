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

	// Offsets depend on the index
	// Second go routine should start at bufferSize
	for i := 0; i < fileParts; i++ {
		chunks[i].size = bufferSize
		chunks[i].offset = int64(bufferSize * i)
	}

	// check for any left over bytes. Add the residual number of bytes as the
	// the last chunk size.
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

	return buffer
}
