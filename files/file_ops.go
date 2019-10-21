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

// LoadContentsOfFile reads the entire contents of the file
// specified by the dataPath parameter.
func LoadContentsOfFile(dataPath string) []byte {
	file, err := os.Open(dataPath)
	checkErr(err)
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	checkErr(err)

	return fileData
}

// WriteToNewFile writes  contents of byte slice to file specified
// by the fileName parameter.
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

// RemoveFile deletes the file specifed by the dataPath parameter.
func RemoveFile(dataPath string) {
	delErr := os.Remove(dataPath)
	checkErr(delErr)
	log.Infof("Original file '%s' has been removed", dataPath)
}
