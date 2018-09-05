package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AESGCMEncrypt(plaintext *[]byte, key *[]byte) ([]byte, error) {
	c, err := aes.NewCipher(*key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, *plaintext, nil), nil
}

func AESGCMDecrypt(ciphertext *[]byte, key *[]byte) ([]byte, error) {
	c, err := aes.NewCipher(*key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(*ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, encrypted := (*ciphertext)[:nonceSize], (*ciphertext)[nonceSize:]
	return gcm.Open(nil, nonce, encrypted, nil)
}
