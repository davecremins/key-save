package files

import (
	"bytes"
	"testing"
)

func TestWritingToFileViaChannel(t *testing.T) {
	var buf bytes.Buffer
	byteChan := make(chan *[]byte, 2)
	arg1, arg2 := []byte("Hello "), []byte("World")
	byteChan <- &arg1
	byteChan <- &arg2
	close(byteChan)
	write(&buf, byteChan, false)
	if "Hello World" != buf.String() {
		t.Error("write via channel returned unexpected result:", buf.String())
	}
}
