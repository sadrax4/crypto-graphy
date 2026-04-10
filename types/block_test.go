package types

import (
	"testing"

	"github.com/sadrax4/crypto-graphy/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	block := utils.RandomBlockGen()
	hashBlock := HashBlock(block)
	assert.Equal(t, 32, len(hashBlock))
}
