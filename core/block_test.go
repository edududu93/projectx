package core

import (
	"testing"
	"time"
	//"fmt"

	"projectx/types"
	"projectx/crypto"

	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:		1,
		PrevBlockHash:  types.RandomHash(),
		Height: 		height,
		Timestamp: 		time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}