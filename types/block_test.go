package types

import (
	"testing"

	"github.com/sadrax4/crypto-graphy/crypto"
	"github.com/sadrax4/crypto-graphy/utils"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	var (
		block      = utils.RandomBlockGen()
		hashBlock  = HashBlock(block)
		privateKey = crypto.GeneratePrivateKey()
		publicKey  = privateKey.Public()
	)
	sig := SignBlock(privateKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(hashBlock, publicKey))
}

func TestHashBlock(t *testing.T) {
	block := utils.RandomBlockGen()
	hashBlock := HashBlock(block)
	assert.Equal(t, 32, len(hashBlock))
}
