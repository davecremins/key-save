package ioops

import (
	"bytes"
	"strings"
	"testing"
)

func TestChunkCalculation(t *testing.T) {
	chunks := PrepareChunks(10089, 128)
	expectedChunkCount := func(Size, bufferSize int) int {
		divideResult := Size / bufferSize
		remainder := Size % bufferSize
		if remainder > 0 {
			divideResult++
		}
		return divideResult
	}(10089, 128)
	if expectedChunkCount != len(*chunks) {
		t.Error("prepareChunks failed to create the correct Size array")
	}
}

func TestChunkingOfReader(t *testing.T) {
	reader := strings.NewReader("Testing my new reader for the files package")

	part1 := Chunk{Size: 7, Offset: int64(0)}
	ReadIntoChunk(reader, &part1)

	expected := []byte("Testing")
	if !bytes.Equal(expected, *(part1.Data)) {
		t.Error("read failed to extract correct data for part 1")
	}

	part2 := Chunk{Size: 18, Offset: int64(7)}
	ReadIntoChunk(reader, &part2)

	expected = []byte(" my new reader for")
	if !bytes.Equal(expected, *(part2.Data)) {
		t.Error("read failed to extract correct data for part 2")
	}
}

func TestPanicInChunking(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("chunked reading should have paniced with EOF error")
		} else {
			t.Log("panic raised:", r)
		}
	}()

	reader := strings.NewReader("Hello World")
	part := Chunk{Size: 12, Offset: int64(0)}
	ReadIntoChunk(reader, &part)
}

func TestReadingOfChunks(t *testing.T) {
	const chunkCount = 3
	data := []byte("5\nFirstHelloWorld")
	chunks := ReadChunks(data)

	if len(chunks) != chunkCount {
		t.Error("reading failed to extract correct amount of chunks")
	}

	chunk := chunks[0]
	expected := []byte("First")

	if !bytes.Equal(expected, *(chunk.Data)) {
		t.Error("read failed to extract correct data")
		t.Log(string(*(chunk.Data)))
	}
}
