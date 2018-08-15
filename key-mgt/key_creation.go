package keymgt

import (
	"errors"
	"crypto/rand"
	"crypto/rsa"
)

var supportedAESKeySizes = map[int]bool{
	128: true,
	192: true,
	256: true,
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

func CreateAESKey(keySize int) (string, error) {
	if !supportedAESKeySizes[keySize] {
		return "", errors.New("AES key size only supports 128, 192 or 256 bits")
	}
	key := make([]byte, keySize/8)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return string(key), nil
}
