package keyMgt

import (
	"bufio"
	"fmt"
	"os"
	//"encoding/base64"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func LoadPublicKeyFromPemFile(file string) (*rsa.PublicKey, error) {
	fmt.Println("Loading key from", file)

	f, _ := os.Open(file)
	defer f.Close()
	fileInfo, _ := f.Stat()
	size := fileInfo.Size()

	pemBytes := make([]byte, size)
	buffer := bufio.NewReader(f)
	_, err := buffer.Read(pemBytes)

	fmt.Println("Pem data as string", string(pemBytes))

	block, _ := pem.Decode([]byte(pemBytes))
	if block == nil {
		return nil, errors.New("fatal: public key decoding error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := pubInterface.(*rsa.PublicKey)

	return publicKey, nil
}

func LoadPrivateKeyFromPemFile(file string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode([]byte(data))
	if block == nil {
		return nil, errors.New("fatal: private key decoding error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
