package keyMgt

import (
	"testing"
	//"fmt"
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

func TestConversionOfPublicKey(t *testing.T){}
func TestConversionOfPrivateKey(t *testing.T){}
