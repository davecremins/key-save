package safe

import (
	cipher "github.com/davecremins/safe-deposit-box/cipher"
	fr "github.com/davecremins/safe-deposit-box/files"
	ops "github.com/davecremins/safe-deposit-box/io-ops"
)

// TODO
func ProtectFile(fileToProtect string, key []byte) err {
	// Open file / defer close
	file, err := fr.Open(fileToProtect)
	if err != nil {
		panic(err)
	}
	defer file.close()

	// Get size of file
	fileSize, err := fr.Size(file)
	if err != nil {
		panic(err)
	}
	// Calculate amount of chunks required
	chunks := ops.PrepareChunks(fileSize, 256)
	chunkAmount := len(*chunks)

	// Create channel to receive byte chunks for encryption
	bytesForEncryptionCh := make(chan *[]byte, chunkAmount)
	go fr.ReadFileInChunks(file, chunks, bytesForEncryptionCh)

	// Begin file read sending chunks to encryption channel
	cipherResultCh := make(chan *[]byte, chunkAmount)
	// Generate AES key
	aesKey, err := cipher.CreateRandomKey(32)
	if err != nil {
		panic(err)
	}
	createWorkers(bytesForEncryptionCh, &aesKey, cipherResultCh, chunkAmount)

	// Once encryption channel is closed

	// Send signal to create file and write encrypted content


}

func createWorkers(bytesToEncrypt chan *[]byte, cipherResult chan *[]byte, workerCount int) {
	var wg sync.WaitGroup
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go encryptWorker(bytesToEncrypt, cipherResult, &wg)
	}
	wg.Wait()
	close(cipherResult)
}

func encryptWorker(bytesToEncrypt <-chan *[]byte, cipherResult chan<- *[]byte, wg *sync.WaitGroup){
	for bytes := range bytesToEncrypt {
		ciphertext, err := cipher.AESGCMEncrypt(bytes, key)
		if err != nil {
			panic(err)
		}
		cipherResult <- ciphertext
	}
	wg.Done()
}
