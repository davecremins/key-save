package keymgt

import "testing"

// TODO: Fix failing test
func TestCanDecodeBytesToPemBlock(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	encodingStruct, _ := createPemEncodingStructureForKey(privateKey)
	block, err := decodeBytesToPemBlock(&encodingStruct.block.Bytes)
	t.Log(block)
	if err != nil {
		t.Error("failed to decode pem block:", err)
	}
}

func TestCanConvertPemBlockToPublicKey(t *testing.T){}
func TestCanConvertPemBlockToPrivateKey(t *testing.T){}
