package apps

import (
	"os"
	"io/ioutil"
	"flag"
	"path/filepath"
	log "github.com/sirupsen/logrus"
	km "gitlab.com/davecremins/safe-deposit-box/key-mgt"
	"gitlab.com/davecremins/safe-deposit-box/cipher"
)

const (
	empty   = "none"
	encrypt = "encrypt"
	decrypt = "decrypt"
)

type CLI struct{}

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
	default:
		log.Fatal("Unsupported operation requested")
	}

	/*if *keyPtr == empty {
		log.Fatal("Requested operation requires a key")
	}*/

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func encryptionProcess(dataPath string) {
	if dataPath == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	fileToBeEncrypted, err := os.Open(dataPath)
	checkErr(err)

	fileData, err := ioutil.ReadAll(fileToBeEncrypted)
	checkErr(err)
	fileToBeEncrypted.Close()

	key, err := km.CreateRandomKeyBytes(24)
	checkErr(err)

	encrypted, err := cipher.AESGCMEncrypt(&fileData, &key)
	checkErr(err)

	encryptedFileName := filepath.Base(dataPath) + ".sdb"
	encryptedFile, err := os.OpenFile(
		encryptedFileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	checkErr(err)

	bytesWritten, err := encryptedFile.Write(encrypted)
	checkErr(err)
	encryptedFile.Close()

	log.Info("Wrote %d bytes.\n", bytesWritten)

	delErr := os.Remove(dataPath)
	checkErr(delErr)

	log.Info("base64 key used during encryption process:", km.ConvertToBase64Str(key))
}
