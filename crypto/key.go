package crypto

import "crypto/ed25519"

const (
	privKeyLen = 64
	pubKeyLen  = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Byte() []byte {
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
