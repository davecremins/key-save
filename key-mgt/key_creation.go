package keymgt

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"io"
)

var supportedKeySizes = map[int]bool{
	16: true,
	24: true,
	32: true,
}

func CreateRSAKeys(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func ConvertPublicKeyToInterface(publicKey *rsa.PublicKey) interface{} {
	var iType interface{}
	iType = publicKey
	return iType
}

func ConvertPrivateKeyToInterface(privateKey *rsa.PrivateKey) interface{} {
	var iType interface{}
	iType = privateKey
	return iType
}

func CreateRandomKeyBytes(keySize int) ([]byte, error) {
	if !supportedKeySizes[keySize] {
		return nil, errors.New("Random key size support for 16, 24 or 32 bytes only")
	}
	key := make([]byte, keySize)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

func ConvertToBase64Str(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}
