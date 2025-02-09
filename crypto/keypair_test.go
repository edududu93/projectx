package crypto

import (
	"testing"
	//"fmt"

	"github.com/stretchr/testify/assert"
)

/*
func TestGeneratePrivateKey (t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	//address := pubKey.Address()

	msg := []byte("Hello World!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	b := sig.Verify(pubKey, msg)

	assert.True(t, b)

}

*/

func TestKeypair_Sign_Verify(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()
	msg := []byte("Hello World!")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(PublicKey, msg))
	
}