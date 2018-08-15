package keymgt

import (
	"reflect"
	"testing"
)

func TestErrorIsReturnedForBadBitSizeInRSAKeyCreation(t *testing.T) {
	keySizes := []int{-716, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 11}
	for _, size := range keySizes {
		_, _, err := CreateRSAKeys(size)
		if err == nil {
			t.Errorf("error should have been returned for rsa key bit size %d", size)
		}
	}
}

func TestErrorIsReturnedForBadBitSizeInAESKeyCreation(t *testing.T) {
	keySizes := []int{23, 127, -45, 199, 254, 312, 500}
	for _, size := range keySizes {
		_, err := CreateAESKey(size)
		if err == nil {
			t.Errorf("error should have been returned for aes key bit size %d", size)
		}
	}

}

func TestConversionOfPrivateKey(t *testing.T) {
	privateKey, _, _ := CreateRSAKeys(1024)
	iType := ConvertPrivateKeyToInterface(privateKey)
	typeIs := reflect.TypeOf(iType).String()
	if typeIs != "*rsa.PrivateKey" {
		t.Errorf("type check for *rsa.PrivateKey failed, got %s instead", typeIs)
	}
}

func TestConversionOfPublicKey(t *testing.T) {
	_, publicKey, _ := CreateRSAKeys(1024)
	iType := ConvertPublicKeyToInterface(publicKey)
	typeIs := reflect.TypeOf(iType).String()
	if typeIs != "*rsa.PublicKey" {
		t.Errorf("type check for *rsa.PublicKey failed, got %s instead", typeIs)
	}
}
