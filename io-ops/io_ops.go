package ioops

import (
	"fmt"
	"io"
)

type Chunk struct {
	Sequence int
	Size   int
	Offset int64
}

func PrepareChunks(blobSize int64, bufferSize int) *[]Chunk {
	Size := int(blobSize)
	parts := Size / bufferSize
	chunks := make([]Chunk, parts)

	for i := 0; i < parts; i++ {
		chunks[i].Sequence = i
		chunks[i].Size = bufferSize
		chunks[i].Offset = int64(bufferSize * i)
	}

	// Add the remaining number of bytes as last Chunk Size
	if remainder := Size % bufferSize; remainder != 0 {
		c := Chunk{Sequence: parts, Size: remainder, Offset: int64(parts * bufferSize)}
		chunks = append(chunks, c)
	}

	return &chunks
}

func Read(handle io.ReaderAt, part Chunk) *[]byte {
	buffer := make([]byte, part.Size)
	_, err := handle.ReadAt(buffer, part.Offset)

	if err != nil {
		if err == io.EOF {
			fmt.Println("fatal: should not have read EOF")
			panic(err)
		}
	}
	return &buffer
}
