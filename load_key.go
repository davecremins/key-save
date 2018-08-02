package keyMgt

import (
	"fmt"
	"io/ioutil"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func LoadPublicFromFile(file string) (*rsa.PublicKey, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("fatal: public key decoding error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return &pub, nil
}
