package ioops

import (
	"fmt"
	"io"
)

type Chunk struct {
	Size   int
	Offset int64
	Data   *[]byte
}

func PrepareChunks(blobSize int64, bufferSize int) *[]Chunk {
	size := int(blobSize)
	parts := size / bufferSize
	chunks := make([]Chunk, parts)

	for i := 0; i < parts; i++ {
		chunks[i].Size = bufferSize
		chunks[i].Offset = int64(bufferSize * i)
	}

	// Add the remaining number of bytes as last Chunk Size
	if remainder := size % bufferSize; remainder != 0 {
		c := Chunk{Size: remainder, Offset: int64(parts * bufferSize)}
		chunks = append(chunks, c)
	}

	return &chunks
}

func ReadIntoChunk(handle io.ReaderAt, part *Chunk) {
	buffer := make([]byte, part.Size)
	_, err := handle.ReadAt(buffer, part.Offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("fatal: should not have read EOF")
			panic(err)
		}
	}
	part.Data = &buffer
}
