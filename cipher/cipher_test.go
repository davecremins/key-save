package cipher

import (
	"bytes"
	"testing"

	km "github.com/davecremins/safe-deposit-box/key-mgt"
)

func TestRSAEncryptionWithPublicKeyIsUnique(t *testing.T) {
	_, publicKey, _ := km.CreateRSAKeys(2048)
	message := []byte("Testing encryption function")
	ciphertext1, _ := RSAEncrypt(&message, publicKey)
	ciphertext2, _ := RSAEncrypt(&message, publicKey)
	if bytes.Equal(ciphertext1, ciphertext2) {
		t.Error("rsa encryption is not unique")
	}
}

func TestRSAEncryptionDecryptionProcess(t *testing.T) {
	privateKey, publicKey, _ := km.CreateRSAKeys(4096)
	message := []byte("Testing full RSA encryption decryption process")
	ciphertext, _ := RSAEncrypt(&message, publicKey)
	plaintext, _ := RSADecrypt(&ciphertext, privateKey)
	if !bytes.Equal(message, plaintext) {
		t.Error("RSA encryption/decryption process failed, plaintext doesn't match original message")
	}
}

func TestAESEncryptionDecryptionProcess(t *testing.T) {
	aesKey, _ := km.CreateRandomKeyBytes(32)
	message := []byte("Testing full AES encryption decryption process")
	ciphertext, _ := AESGCMEncrypt(&message, &aesKey)
	plaintext, _ := AESGCMDecrypt(&ciphertext, &aesKey)
	if !bytes.Equal(message, plaintext) {
		t.Error("AES encryption/decryption process failed, plaintext doesn't match original message")
	}
}
