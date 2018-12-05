package keymgt

import (
	"bytes"
	"testing"
)

func TestCanDecodePemEncodedKeyToPemBlock(t *testing.T) {
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

func TestCanConvertPemBlockToPublicKey(t *testing.T) {
	_, publicKey, _ := CreateRSAKeys(2048)
	out := new(bytes.Buffer)
	PemEncodeKeyToOutput(publicKey, out)
	key := LoadPublicKeyFromPemData(out)
	if key == nil {
		t.Error("failed to load public key")
	}
}

func TestCanConvertPemBlockToPrivateKey(t *testing.T) {
	privateKey, _, _ := CreateRSAKeys(2048)
	out := new(bytes.Buffer)
	PemEncodeKeyToOutput(privateKey, out)
	_, err := LoadPrivateKeyFromPemData(out)
	if err != nil {
		t.Error("failed to load private key")
	}
}
