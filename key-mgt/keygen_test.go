package keymgt

import (
	"reflect"
	"testing"
)

func TestCreationOfRSAKeys(t *testing.T) {
	privateKey, publicKey, err := CreateRSAKeys(1024)
	if privateKey == nil {
		t.Error("failed to create privateKey")
	}

	if publicKey == nil {
		t.Error("failed to create publicKey")
	}

	if err != nil {
		t.Error(err)
	}
}

func TestErrorIsReturnedForBadBitSizeInRSAKeyCreation(t *testing.T) {
	keySizes := []int{-716, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 11}
	for _, size := range keySizes {
		_, _, err := CreateRSAKeys(size)
		if err == nil {
			t.Errorf("error should have been returned for rsa key bit size %d", size)
		}
	}
}

func TestErrorIsReturnedForBadByteSizeInRandomKeyCreation(t *testing.T) {
	keySizes := []int{0, 1, 15, 21, 29, 33, 45, 64, 82, 128, 256, 512, 1024}
	for _, size := range keySizes {
		_, err := CreateRandomKeyBytes(size)
		if err == nil {
			t.Errorf("error should have been returned for key byte size %d", size)
		}
	}

}

func TestCreateRandomKey(t *testing.T) {
	key, _ := CreateRandomKeyBytes(32)
	t.Log(key)
}

func TestConvertRandomByteKeyToBase64Str(t *testing.T) {
	key, _ := CreateRandomKeyBytes(32)
	base64Str := ConvertToBase64Str(key)
	t.Log(base64Str)
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
