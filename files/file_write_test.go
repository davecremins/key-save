package files

import (
	"bytes"
	"strings"
	"testing"
)

func TestWritingToFileViaChannel(t *testing.T) {
	var buf bytes.Buffer
	byteChan := make(chan *[]byte)
	go func(ch chan *[]byte) {
		arg1, arg2 := []byte("Hello "), []byte("World")
		ch <- &arg1
		ch <- &arg2
		close(ch)

	}(byteChan)
	write(&buf, byteChan, false)
	if "Hello World" != buf.String() {
		t.Error("buffer string contents does not match expected:", buf.String())
	}
}

func TestReadingOfContentInChunks(t *testing.T) {
	reader := strings.NewReader("Testing my new reader for the files package")
	readInChunks(reader, reader.Size(), 16)
}
