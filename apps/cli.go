package apps

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gitlab.com/davecremins/safe-deposit-box/cipher"
	"gitlab.com/davecremins/safe-deposit-box/files"
	km "gitlab.com/davecremins/safe-deposit-box/key-mgt"
	"path/filepath"
	"strings"
)

const (
	empty     = "none"
	encrypt   = "encrypt"
	decrypt   = "decrypt"
	keygen    = "keygen"
	extension = ".sdb"
)

// CLI is an implementation of a command line interface.
type CLI struct{}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run is an interface implemented by CLI that provides the capability
// to encrypt and decrypt data.
func (c *CLI) Run() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Starting CLI")

	var operation, dataPath, key string

	flag.StringVar(&operation, "o", "encrypt", "Operations - encrypt|decrypt|keygen")
	flag.StringVar(&key, "k", empty, "Security key")
	flag.StringVar(&dataPath, "p", empty, "Path to file containing data")

	flag.Parse()

	switch operation {
	case encrypt:
		log.Info("Encryption option requested")
		encryptionProcess(dataPath)
	case decrypt:
		log.Info("Decryption operation requested")
		decryptionProcess(dataPath, key)
	case keygen:
		log.Info("Key generation operation requested")
		keyGenerationProcess()
	default:
		log.Fatal("Unsupported operation requested, please use -h flag for help")
	}

}

func encryptionProcess(dataPath string) {
	if dataPath == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	fileData := files.LoadContentsOfFile(dataPath)
	key, err := km.CreateRandomKeyBytes(24)
	checkErr(err)

	encrypted, err := cipher.AESGCMEncrypt(&fileData, &key)
	checkErr(err)

	encryptedFileName := filepath.Base(dataPath) + extension
	files.WriteToNewFile(encryptedFileName, encrypted)

	files.RemoveFile(dataPath)

	log.Info("base64 key used during encryption process: ", km.ConvertToBase64Str(key))
}

func decryptionProcess(dataPath string, key string) {
	if dataPath == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	if key == empty {
		log.Fatal("Requested operation requires a key")
	}

	fileData := files.LoadContentsOfFile(dataPath)
	byteKey, err := km.ConvertBase64StrToBytes(key)
	checkErr(err)

	plaintext, err := cipher.AESGCMDecrypt(&fileData, &byteKey)
	checkErr(err)

	decryptedFileName := filepath.Base(dataPath)
	decryptedFileName = strings.TrimRight(decryptedFileName, filepath.Ext(decryptedFileName))
	files.WriteToNewFile(decryptedFileName, plaintext)

	files.RemoveFile(dataPath)
}

func keyGenerationProcess() {
	key, err := km.CreateRandomKeyBytes(24)
	checkErr(err)
	log.Info("Key created (base64): ", km.ConvertToBase64Str(key))
}
