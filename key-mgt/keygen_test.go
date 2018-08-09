package keymgt

import (
	"reflect"
	"testing"
)

func TestErrorIsReturnedForBadBitSize(t *testing.T) {
	bitSizes := []int{-716, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 11}
	for _, size := range bitSizes {
		_, _, err := CreateRSAKeys(size)
		if err == nil {
			t.Errorf("error should have been returned for rsa key bit size %d", size)
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
