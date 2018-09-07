package files

import (
	"bytes"
	"strings"
	"testing"
)

func TestChunkCalculation(t *testing.T) {
	chunks := prepareChunks(10089, 128)
	expectedChunkCount := func(size, bufferSize int) int {
		divideResult := size / bufferSize
		remainder := size % bufferSize
		if remainder > 0 {
			divideResult++
		}
		return divideResult
	}(10089, 128)
	if expectedChunkCount != len(*chunks) {
		t.Error("prepareChunks failed to create the correct size array")
	}
}

func TestChunkingOfReader(t *testing.T) {
	reader := strings.NewReader("Testing my new reader for the files package")

	part1 := chunk{size: 7, offset: int64(0)}
	buffer := read(reader, part1)

	expected := []byte("Testing")
	if !bytes.Equal(expected, buffer) {
		t.Error("read failed to extract correct data for part 1")
	}

	part2 := chunk{size: 18, offset: int64(7)}
	buffer = read(reader, part2)

	expected = []byte(" my new reader for")
	if !bytes.Equal(expected, buffer) {
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
	part := chunk{size: 12, offset: int64(0)}
	read(reader, part)
}
