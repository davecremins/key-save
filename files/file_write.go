package files

import (
	"os"
)

const separator = "|"

func CreateAndWaitToWrite(filename string, bytesToWrite chan *[]byte, includeSeparator bool) {
	handle, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	for content := range bytesToWrite {
		handle.Write(*content)
		if includeSeparator {
			handle.WriteString(separator)
		}
	}
	handle.Sync()
}
