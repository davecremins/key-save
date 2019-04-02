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

// CreateRSAKeys creates a private/public RSA key pair.
func CreateRSAKeys(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// ConvertPublicKeyToInterface creates an interface abstraction for an RSA
// public key.
func ConvertPublicKeyToInterface(publicKey *rsa.PublicKey) interface{} {
	var iType interface{}
	iType = publicKey
	return iType
}

// ConvertPrivateKeyToInterface creates an interface abstraction for an RSA
// private key.
func ConvertPrivateKeyToInterface(privateKey *rsa.PrivateKey) interface{} {
	var iType interface{}
	iType = privateKey
	return iType
}

// CreateRandomKeyBytes creates a random slice of bytes with support for
// 16, 24 and 32 byte lengths.
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

// ConvertToBase64Str converts bytes to a base64 encoded string.
func ConvertToBase64Str(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

// ConvertBase64StrToBytes decodes a base64 encoded string and converts to bytes.
func ConvertBase64StrToBytes(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(key)
}
