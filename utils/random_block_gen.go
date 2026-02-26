package utils

import (
	randc "crypto/rand"
	"io"
	"math/rand"
	"time"

	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
)

func RandomHashGen() []byte {
	hash := make([]byte, 32)
	io.ReadFull(randc.Reader, hash)
	return hash
}

func RandomBlockGen() *proto.Block {
	HeaderData := &proto.Header{
		Version:   1,
		Height:    int32(rand.Intn(1000) + 1),
		PrevHash:  RandomHashGen(),
		RootHash:  RandomHashGen(),
		Timestamp: time.Now().UnixNano(),
	}
	return &proto.Block{
		Header: HeaderData,
	}
}
