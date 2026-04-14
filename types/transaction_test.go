package types

import (
	"testing"

	"github.com/sadrax4/crypto-graphy/crypto"
	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
	"github.com/sadrax4/crypto-graphy/utils"
)

func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()

	fromAddress := fromPrivKey.Public().Address().Bytes()
	toAddress := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   utils.RandomHashGen(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}

	ouput1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}
}
