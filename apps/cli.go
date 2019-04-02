package apps

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gitlab.com/davecremins/safe-deposit-box/cipher"
	km "gitlab.com/davecremins/safe-deposit-box/key-mgt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	empty     = "none"
	encrypt   = "encrypt"
	decrypt   = "decrypt"
	extension = ".sdb"
)

// CLI is an implementation of a command line interface.
type CLI struct{}

// Run is an interface implemented by CLI that provides the capability
// to encrypt and decrypt data.
func (c *CLI) Run() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Starting CLI")

	var operation, dataPath, key string

	flag.StringVar(&operation, "op", "encrypt", "Operation to perform - encrypt|decrypt")
	flag.StringVar(&key, "key", empty, "Security key")
	flag.StringVar(&dataPath, "datapath", empty, "Path to file containing data")

	flag.Parse()

	switch operation {
	case encrypt:
		log.Info("Encryption option requested")
		encryptionProcess(dataPath)
	case decrypt:
		log.Info("Decryption operation requested")
		decryptionProcess(dataPath, key)
	default:
		log.Fatal("Unsupported operation requested")
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadContentsOfFile(dataPath string) []byte {
	file, err := os.Open(dataPath)
	checkErr(err)
	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	checkErr(err)

	return fileData
}

func writeToNewFile(fileName string, content []byte) {
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

func removeFile(dataPath string) {
	delErr := os.Remove(dataPath)
	checkErr(delErr)
	log.Infof("Original file '%s' has been removed", dataPath)
}

func encryptionProcess(dataPath string) {
	if dataPath == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	fileData := loadContentsOfFile(dataPath)
	key, err := km.CreateRandomKeyBytes(24)
	checkErr(err)

	encrypted, err := cipher.AESGCMEncrypt(&fileData, &key)
	checkErr(err)

	encryptedFileName := filepath.Base(dataPath) + extension
	writeToNewFile(encryptedFileName, encrypted)

	removeFile(dataPath)

	log.Info("base64 key used during encryption process: ", km.ConvertToBase64Str(key))
}

func decryptionProcess(dataPath string, key string) {
	if dataPath == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	if key == empty {
		log.Fatal("Requested operation requires a key")
	}

	fileData := loadContentsOfFile(dataPath)
	byteKey, err := km.ConvertBase64StrToBytes(key)
	checkErr(err)

	plaintext, err := cipher.AESGCMDecrypt(&fileData, &byteKey)
	checkErr(err)

	decryptedFileName := filepath.Base(dataPath)
	decryptedFileName = strings.TrimRight(decryptedFileName, filepath.Ext(decryptedFileName))
	writeToNewFile(decryptedFileName, plaintext)

	removeFile(dataPath)
}
