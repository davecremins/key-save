package keymgt

import (
	"crypto/rsa"
	"bytes"
	"crypto/x509"
	"errors"
	"encoding/pem"
	"io"
)

func LoadPublicKeyFromPemData(reader io.Reader) (*rsa.PublicKey, error) {
	pemBytes := loadDataFromSource(reader)
	block, err := decodeBytesToPemBlock(&pemBytes)
	if err != nil {
		panic(err)
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := pubInterface.(*rsa.PublicKey)
	return publicKey, nil
}

func LoadPrivateKeyFromPemData(reader io.Reader) (*rsa.PrivateKey, error) {
	pemBytes := loadDataFromSource(reader)
	block, err := decodeBytesToPemBlock(&pemBytes)
	if err != nil {
		panic(err)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
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
