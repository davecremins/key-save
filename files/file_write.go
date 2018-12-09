package files

import "io"

func WaitToWrite(out io.Writer, bytesToWrite chan *[]byte, finished chan bool) {
	go func() {
		for content := range bytesToWrite {
			out.Write(*content)
		}
		finished <- true
	}()
}
