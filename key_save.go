package keyMgt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type keyEncoding struct {
	block   *pem.Block
	keyType string
}

func CreateRSAKey(keySize int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	return privateKey, &privateKey.PublicKey, nil
}

//Maybe support file location
func CreateFile(key interface{}) string {
	keyEncodingData, err := pemBlockForKey(key)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	fileName := "rsa" + keyEncodingData.keyType + ".pem"
	keyOut, err := os.Create(fileName)
	if err != nil {
		fmt.Println("failed to open %s for writing:", fileName, err)
		os.Exit(1)
	}
	defer keyOut.Close()
	pem.Encode(keyOut, keyEncodingData.block)
	return fileName
}

func pemBlockForKey(key interface{}) (*keyEncoding, error) {
	switch k := key.(type) {
	case *rsa.PublicKey:
		return &keyEncoding{
			&pem.Block{Type: "BEGIN RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(k)},
			"_public",
		}, nil
	case *rsa.PrivateKey:
		return &keyEncoding{
			&pem.Block{Type: "BEGIN RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)},
			"_private",
		}, nil
	default:
		return nil, fmt.Errorf("Unsupported key type %s", k)
	}
}

//verifier
/*func verifySignatureWithPublicKey(message string, signature []byte, key *rsa.PublicKey) {
	newhash := crypto.SHA256
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	pssh := newhash.New()
	pssh.Write([]byte(message))
	hashed := pssh.Sum(nil)
	err := rsa.VerifyPSS(
		key,
		newhash,
		hashed,
		signature,
		&opts)
	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify Signature successful")
	}
}

//decrypter
func getPlainTextWithPrivateKey(ciphertext []byte, key *rsa.PrivateKey) []byte {
	hash := sha256.New()
	label := []byte("")
	plainText, err := rsa.DecryptOAEP(
		hash,
		rand.Reader,
		key,
		ciphertext,
		label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return plainText
}

//signer
func getSignatureWithPrivKey(message string, key *rsa.PrivateKey) []byte {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write([]byte(PSSmessage))
	hashed := pssh.Sum(nil)
	signature, err := rsa.SignPSS(
		rand.Reader,
		key,
		newhash,
		hashed,
		&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return signature
}

//encrypter
func getCypherTextWithPubKey(msg string, key *rsa.PublicKey) []byte {
	message := []byte(msg)
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		key,
		message,
		label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return ciphertext
}*/
