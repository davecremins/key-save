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
	// TODO: Needs a test for the channel usage
	chunkChannel := make(chan *[]byte, chunkAmount)

	for i := 0; i < chunkAmount; i++ {
		go read(file, (*chunks)[i], chunkChannel)
	}

	for bytesRead := range chunkChannel {
		fmt.Println("Bytes read:", string(*bytesRead))
	}
	close(chunkChannel)
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
func read(handle io.ReaderAt, part chunk, chunkOut chan<- *[]byte) {
	buffer := make([]byte, part.size)
	_, err := handle.ReadAt(buffer, part.offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("fatal: should not have read EOF")
			panic(err)
		}
	}
	chunkOut <- &buffer
}
