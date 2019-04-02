package ioops

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const delimiter = '\n'

type pChunkSlice *[]Chunk

// Chunk represents a piece of data, the size of the data
// and its offset within its origin.
type Chunk struct {
	Size   int
	Offset int64
	Data   *[]byte
}

// PrepareChunks calculates the amount of chunks and their offsets based on
// the originating data's size and the requested buffer size.
func PrepareChunks(blobSize int64, bufferSize int) *[]Chunk {
	size := int(blobSize)
	parts := size / bufferSize
	chunks := make([]Chunk, parts)

	for i := 0; i < parts; i++ {
		chunks[i].Size = bufferSize
		chunks[i].Offset = int64(bufferSize * i)
	}

	// Add the remaining number of bytes as last chunk size
	if remainder := size % bufferSize; remainder != 0 {
		c := Chunk{Size: remainder, Offset: int64(parts * bufferSize)}
		chunks = append(chunks, c)
	}

	return &chunks
}

// ReadIntoChunk reads the data from the reader into the chunk
// based on the chunks pre-calculated offset.
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

func extractChunkSize(reader io.Reader) (size, firstChunkPos int, err error) {
	br := bufio.NewReader(reader)
	chunkNumberStr, err := br.ReadString(delimiter)
	if err != nil {
		return 0, 0, err
	}
	firstChunkPos = len(chunkNumberStr)

	// Remove delimiter
	chunkNumberStr = strings.TrimSuffix(chunkNumberStr, string(delimiter))
	chunkSize, err := strconv.Atoi(chunkNumberStr)
	if err != nil {
		panic(errors.New("fatal: unable to convert chunk number to int"))
	}
	return chunkSize, firstChunkPos, nil
}

// ReadChunks takes a byte sequence and extracts the chunk size,
// moves the reading position to the first chunk in the byte sequence
// and reads all of the chunks.
func ReadChunks(b []byte) []Chunk {
	reader := bytes.NewReader(b)
	chunkSize, firstChunkPos, err := extractChunkSize(reader)
	if err != nil {
		panic(errors.New("fatal: unable to extract chunk size"))
	}

	reader.Seek(int64(firstChunkPos), 0)
	data := make([]byte, chunkSize)
	chunks := []Chunk{}

	for {
		n, err := reader.Read(data)
		if err == io.EOF {
			break
		}
		dataChunk := make([]byte, n)
		copy(dataChunk, data[:n])
		c := Chunk{Size: n, Data: &dataChunk}
		chunks = append(chunks, c)
	}

	return chunks
}
