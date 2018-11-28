package keymgt

import "testing"

/*func TestCanDecodeBytesToPemBlock(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	encodingStruct, _ := pemBlockForKey(privateKey)
	_, err := decodeBytesToPemBlock(&encodingStruct.block.Bytes)
	if err != nil {
		t.Error("failed to decode pem block:", err)
	}
}*/

func TestCanConvertPemBlockToPublicKey(t *testing.T){}
func TestCanConvertPemBlockToPrivateKey(t *testing.T){}
