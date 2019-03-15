package apps

import (
	"os"
	"io/ioutil"
	"flag"
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

	operationPtr := flag.String("op", "encrypt", "Operation to perform - encrypt|decrypt")
	//keyPtr := flag.String("key", empty, "Security key")
	dataLocationPtr := flag.String("datapath", empty, "Path to file containing data")

	flag.Parse()

	switch *operationPtr {
	case encrypt:
		log.Info("Encryption operation requested")
	case decrypt:
		log.Info("Decryption operation requested")
	default:
		log.Fatal("Unsupported operation requested")
	}

	/*if *keyPtr == empty {
		log.Fatal("Requested operation requires a key")
	}*/

	if *dataLocationPtr == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}

	// Open file from data location
	file, err := os.Open(*dataLocationPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read in file
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Generate random key
	key, err := km.CreateRandomKeyBytes(24)
	if err != nil {
		log.Fatal(err)
	}

	// encrypt the file
	encrypted, err := cipher.AESGCMEncrypt(&fileData, &key)
	if err != nil {
		log.Fatal(err)
	}

	// write to new file
	encryptedFile, err := os.OpenFile(
		"encrypted.data",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer encryptedFile.Close()

	// Write bytes to file
	bytesWritten, err := encryptedFile.Write(encrypted)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Wrote %d bytes.\n", bytesWritten)

	// delete original

	// show key to user
	log.Info("Key used during encryption process:", string(key[:])) 
}
