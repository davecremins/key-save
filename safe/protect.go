package safe

/*import (
	cipher "gitlab.com/davecremins/safe-deposit-box/cipher"
	fr "gitlab.com/davecremins/safe-deposit-box/files"
)

// TODO
func ProtectFile(fileToProtect string, key []byte) err {
	// Open file / defer close

	// Get size of file

	// Calculate amount of chunks required

	// Create channel to receive byte chunks for encryption

	// Allocate workers for encryption process

	// Begin file read sending chunks to encryption channel
	// Once encryption channel receives it can begin encryption

	// Once encryption channel is closed

	// Send signal to create file and write encrypted content

	// Generate AES key
	aesKey, err := cipher.CreateRandomKey(32)
	if err != nil {
		panic(err)
	}

	bytesReadCh := (<-chan *[]byte)
	fr.ReadFileInChunks(fileToProtect, 128)
}*/
