package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func RSAEncrypt(content *[]byte, publicKey *rsa.PublicKey){
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, *content, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(*content), ciphertext)
	fmt.Println()
}

func RSADecrypt(ciphertext *[]byte, privateKey *rsa.PrivateKey){
	label := []byte("")
	hash := sha256.New()

	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, *ciphertext, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)
	fmt.Println()
}
