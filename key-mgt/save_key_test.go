package keymgt

import "testing"

func TestCreatePemBlockForPrivateKey(t *testing.T){
	privateKey, _, _ := CreateRSAKeys(1024)
	encodingStruct, err := pemBlockForKey(privateKey)
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
	encodingStruct, err := pemBlockForKey(publicKey)
	if err != nil {
		t.Error(err)
	}

	if encodingStruct.block == nil {
		t.Error("failed to create PEM block for public key")
	}

	t.Log(*encodingStruct.block)
}

func TestErrorIsReturnedForUnsupportedType(t *testing.T){
	_, err := pemBlockForKey("key")
	if err == nil {
		t.Error("pemBlockForKey should have failed with unsuported type")
	}
	t.Log(err)
}

// Does this belong here?
func TestCreateFileNameConvention(t *testing.T){
	fileName := createFileName(&keyEncoding{keyType: "_public"})
	if fileName != "rsa_public.pem" {
		t.Error("Incorrect file name convention")
	}
}
