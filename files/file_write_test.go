package files

import (
	"bytes"
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
		t.Error("write via channel returned unexpected result:", buf.String())
	}
}
