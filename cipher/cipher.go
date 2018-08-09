package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func RSAEncrypt(content *[]byte, publicKey *rsa.PublicKey) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, *content, label)

	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func RSADecrypt(ciphertext *[]byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()

	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, *ciphertext, label)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
