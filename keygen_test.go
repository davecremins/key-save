package keyMgt

import (
	"testing"
	"fmt"
)

func TestBadBitSize(t *testing.T) {
	bitSize := -1
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RSA key gen panic for key bit size", bitSize)
		} else {
			t.Errorf("RSA key gen should panic for key bit size %d", bitSize)
		}
	}()
	CreateRSAKeys(bitSize)
}
