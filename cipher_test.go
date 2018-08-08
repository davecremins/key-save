package cipher

import (
	"bytes"
	"testing"
)

func TestEncryptionWithPublicKeyIsUnique(t *testing.T) {
	_, publicKey, _ := CreateRSAKeys(2048)
	message := []byte("Testing encryption function")
	ciphertext1, _ := RSAEncrypt(&message, publicKey)
	ciphertext2, _ := RSAEncrypt(&message, publicKey)
	if bytes.Equal(ciphertext1, ciphertext2) {
		t.Error("rsa encryption is not unique")
	}
}

func TestEncryptionDecryptionProcess(t *testing.T) {
	privateKey, publicKey, _ := CreateRSAKeys(4096)
	message := []byte("Testing full encryption decryption process")
	ciphertext, _ := RSAEncrypt(&message, publicKey)
	plainText, _ := RSADecrypt(&ciphertext, privateKey)
	if !bytes.Equal(message, plainText) {
		t.Error("rsa encryption/decryption process failed, plaintext doesn't match original message")
	}
}
