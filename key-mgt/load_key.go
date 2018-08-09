package key-mgt

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func LoadPublicKeyFromPemFile(fileName string) (*rsa.PublicKey, error) {
	pemBytes := loadKeyFromFile(fileName)
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

func LoadPrivateKeyFromPemFile(fileName string) (*rsa.PrivateKey, error) {
	pemBytes := loadKeyFromFile(fileName)
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

func loadKeyFromFile(fileName string) []byte {
	fmt.Println("Will load key from", fileName)
	pemFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer pemFile.Close()

	pfInfo, _ := pemFile.Stat()
	size := pfInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(pemFile)
	buffer.Read(pemBytes)

	fmt.Println("Pem data read successfully.")
	fmt.Println(string(pemBytes))
	return pemBytes
}

func decodeBytesToPemBlock(pemBytes *[]byte) (*pem.Block, error) {
	block, _ := pem.Decode(*pemBytes)
	if block == nil {
		return nil, errors.New("fatal: key decoding error")
	}
	return block, nil
}
