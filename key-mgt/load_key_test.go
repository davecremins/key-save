package keymgt

import (
	"testing"
	"bytes"
)

func TestCanDecodePemEncodedKeyToPemBlock(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	out := new(bytes.Buffer)
	PemEncodeKeyToOutput(privateKey, out)
	pemEncodedBytes := out.Bytes()
	block, err := decodeBytesToPemBlock(&pemEncodedBytes)
	if err != nil {
		t.Error("failed to decode pem block:", err)
	}
	t.Log(block)
}

func TestCanConvertPemBlockToPublicKey(t *testing.T){}
func TestCanConvertPemBlockToPrivateKey(t *testing.T){}
