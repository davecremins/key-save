package keymgt

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

// LoadPublicKeyFromPemData extracts the pem encoded public key
// and converts it to a public key structure.
func LoadPublicKeyFromPemData(reader io.Reader) *rsa.PublicKey {
	pemBytes := loadDataFromSource(reader)
	block, err := decodeBytesToPemBlock(&pemBytes)
	if err != nil {
		panic(err)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("FATAL: failed to parse DER encoded public key: " + err.Error())
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	return publicKey
}

// LoadPrivateKeyFromPemData extracts the pem encoded private key
// and converts it to a private key structure.
func LoadPrivateKeyFromPemData(reader io.Reader) (*rsa.PrivateKey, error) {
	pemBytes := loadDataFromSource(reader)
	block, err := decodeBytesToPemBlock(&pemBytes)
	if err != nil {
		panic(err)
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadDataFromSource(reader io.Reader) []byte {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)
	return buffer.Bytes()
}

func decodeBytesToPemBlock(pemBytes *[]byte) (*pem.Block, error) {
	block, _ := pem.Decode(*pemBytes)
	if block == nil {
		return nil, errors.New("FATAL: could not decode pem bytes to block")
	}
	return block, nil
}
