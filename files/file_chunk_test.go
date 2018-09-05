package files

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func TestChunkingOfReader(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	reader := strings.NewReader("Testing my new reader for the files package")

	part1 := chunk{size: 7, offset: int64(0)}
	buffer := read(reader, part1, wg)

	expected := []byte("Testing")
	if !bytes.Equal(expected, buffer) {
		t.Error("Read failed to extract correct data for part 1")
	}

	part2 := chunk{size: 18, offset: int64(7)}
	buffer = read(reader, part2, wg)

	expected = []byte(" my new reader for")
	if !bytes.Equal(expected, buffer) {
		t.Error("Read failed to extract correct data for part 2")
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

	var wg sync.WaitGroup
	wg.Add(1)
	reader := strings.NewReader("Hello World")

	part := chunk{size: 12, offset: int64(0)}
	read(reader, part, wg)
}
