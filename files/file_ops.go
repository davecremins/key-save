package files

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Load entire contents of file specified by the dataPath parameter
func LoadContentsOfFile(dataPath string) []byte {
	file, err := os.Open(dataPath)
	checkErr(err)
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	checkErr(err)

	return fileData
}

// Write contents of byte slice to file specified by fileName
func WriteToNewFile(fileName string, content []byte) {
	file, err := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	checkErr(err)
	defer file.Close()

	bytesWritten, err := file.Write(content)
	checkErr(err)

	log.Infof("Wrote %d bytes.\n", bytesWritten)
}

// Remove file specifed by dataPath parameter
func RemoveFile(dataPath string) {
	delErr := os.Remove(dataPath)
	checkErr(delErr)
	log.Infof("Original file '%s' has been removed", dataPath)
}
