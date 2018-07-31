package keyMgt

import (
	"testing"
	//"fmt"
	"reflect"
)

func TestBadBitSize(t *testing.T) {
	bitSizes := [5]int{-1, 0, -716}
	for _, size := range bitSizes {
		_, _, err := CreateRSAKeys(size)
		if err == nil{
			t.Errorf("error occured for rsa bit size %d: %s", size, err)
		}
	}
}

func TestConversionOfPrivateKey(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	iType := ConvertPrivateKeyToInterface(privateKey)
	typeIs := reflect.TypeOf(iType).String()
	if typeIs != "*rsa.PrivateKey" {
		t.Errorf("type check for *rsa.PrivateKey failed, got %s instead", typeIs)
	}
}

func TestConversionOfPublicKey(t *testing.T){
	_, publicKey, _ := CreateRSAKeys(1024)
	iType := ConvertPublicKeyToInterface(publicKey)
	typeIs := reflect.TypeOf(iType).String()
	if typeIs != "*rsa.PublicKey" {
		t.Errorf("type check for *rsa.PublicKey failed, got %s instead", typeIs)
	}
}
