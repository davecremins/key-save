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
	finished := make(chan bool)
	WaitToWrite(&buf, byteChan, finished)
	<-finished
	if "Hello World" != buf.String() {
		t.Error("buffer string contents does not match expected:", buf.String())
	}
}

func TestReadingOfContentInChunks(t *testing.T) {
	reader := strings.NewReader("Testing my new reader for the files package")
	chunks := ReadInChunks(reader, reader.Size(), 4)
	if "Test" != string(*(*chunks)[0].Data) {
		t.Error("First chunk doesn't contain expected output")
	}

	if "age" != string(*(*chunks)[10].Data) {
		t.Error("Last chunk doesn't contain expected output")
	}
}
