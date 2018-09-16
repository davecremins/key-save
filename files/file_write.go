package files

import (
	"io"
	"os"
)

const separator = "|"

func CreateAndWaitToWrite(filename string, bytesToWrite chan *[]byte, includeSeparator bool) {
	handle, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer handle.Close()
	write(handle, bytesToWrite, includeSeparator)
	handle.Sync()
}

// TODO: Write test
func write(handle io.Writer, bytesToWrite chan *[]byte, includeSeparator bool) {
	for content := range bytesToWrite {
		handle.Write(*content)
		if includeSeparator {
			handle.Write([]byte(separator))
		}
	}
}
