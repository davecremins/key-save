package keymgt

import "testing"

func TestCreatePemBlockForPrivateKey(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	encodingStruct, err := createPemEncodingStructureForKey(privateKey)
	if err != nil {
		t.Error(err)
	}

	if encodingStruct.block == nil {
		t.Error("failed to create PEM block for private key")
	}

	t.Log(*encodingStruct.block)
}

func TestCreatePemBlockForPublicKey(t *testing.T){
	_, publicKey, _ := CreateRSAKeys(1024)
	encodingStruct, err := createPemEncodingStructureForKey(publicKey)
	if err != nil {
		t.Error(err)
	}

	if encodingStruct.block == nil {
		t.Error("failed to create PEM block for public key")
	}

	t.Log(*encodingStruct.block)
}

func TestPemEncodingErrorIsReturnedForUnsupportedType(t *testing.T){
	_, err := createPemEncodingStructureForKey("key")
	if err == nil {
		t.Error("should have failed with unsuported type")
	}
	t.Log(err)
}

func TestCreateNameConvention(t *testing.T){
	name := createPemName(&pemEncoding{keyType: "_public"})
	if name != "rsa_public.pem" {
		t.Error("Incorrect pem naming convention")
	}
}
