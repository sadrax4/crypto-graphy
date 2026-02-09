package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	privKeyLen = 64
	pubKeyLen  = 32
	seedKeyLen = 32
	addressLen = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func NewPrivateKeyFromString(p string) *PrivateKey {
	b, err := hex.DecodeString(p)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedKeyLen {
		panic("invalid private key, must be 32")
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
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

func (p *PublicKey) Address() *Address {
	return &Address{
		value: p.key[len(p.key)-(addressLen):],
	}
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

type Signature struct {
	value []byte
}

func (s *Signature) Verify(msg []byte, pubkey *PublicKey) bool {
	return ed25519.Verify(pubkey.key, msg, s.value)
}

type Address struct {
	value []byte
}

func (a *Address) String() string {
	return hex.EncodeToString(a.value)
}

func (a *Address) Bytes() []byte {
	return a.value
}
