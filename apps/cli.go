package apps

import (
	"flag"
	log "github.com/sirupsen/logrus"
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
	keyPtr := flag.String("key", empty, "Security key")
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

	if *keyPtr == empty {
		log.Fatal("Requested operation requires a key")
	}

	if *dataLocationPtr == empty {
		log.Fatal("Requested operation requires a path to the file containing the data")
	}
}
