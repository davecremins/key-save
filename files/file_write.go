package files

import "io"

// WaitToWrite is a goroutine that will write bytes to an io.Writer
// and signal that it is finished when the bytes channel is closed.
func WaitToWrite(out io.Writer, bytesToWrite chan *[]byte, finished chan bool) {
	go func() {
		for content := range bytesToWrite {
			out.Write(*content)
		}
		finished <- true
	}()
}
