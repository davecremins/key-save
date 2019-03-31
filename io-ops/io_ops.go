package ioops

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"strconv"
	"bytes"
	"errors"
)

const Delimiter = '\n'

type pChunkSlice *[]Chunk

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

func reverseChunks(chunks pChunkSlice) {
	chunkLength := len(*chunks)
	for i, j := 0, chunkLength-1; i<j; i, j = i+1, j-1 {
		(*chunks)[i], (*chunks)[j] = (*chunks)[j], (*chunks)[i]
	}
}

func extractChunkSize(reader io.Reader) (sizeStr string, firstChunkPos int, err error) {
	br := bufio.NewReader(reader)
	chunkSize, err := br.ReadString(Delimiter)
	if err != nil {
		return "", 0, err
	}
	return chunkSize, len(chunkSize), nil
}

func ReadChunks(b []byte) []Chunk {
	reader := bytes.NewReader(b)
	chunkSizeStr, firstChunkPos, err := extractChunkSize(reader)
	if err != nil {
		panic(errors.New("fatal: unable to extract chunk size"))
	}

	reader.Seek(int64(firstChunkPos), 0)
	chunkSizeStr = strings.TrimSuffix(chunkSizeStr, string(Delimiter))
	blockSize, err := strconv.Atoi(chunkSizeStr)
	if err != nil {
		panic(errors.New("fatal: unable to convert chunk size to int"))
	}

	data := make([]byte, blockSize)
	chunks := []Chunk{}

	for {
		n, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		dataChunk := data[:n]
		c := Chunk{Size: n, Data: &dataChunk}
		chunks = append(chunks, c)
	}

	//reverseChunks(&chunks)

	return chunks
}
