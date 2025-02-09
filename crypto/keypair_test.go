package crypto

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey (t *testing.T) {
	privKey := GeneratePrivateKey()
	//pubKey := privKey.PublicKey()
	//address := pubKey.Address()

	msg := []byte("Hello World!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	fmt.Println(sig)
}