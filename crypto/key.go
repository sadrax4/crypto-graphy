package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

const (
	privKeyLen = 64
	pubKeyLen  = 32
	seedKeyLen = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PrivateKey) Public() *PublicKey {
	pKey := make([]byte, pubKeyLen)
	copy(pKey, p.key[32:])
	return &PublicKey{
		key: pKey,
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedKeyLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}
