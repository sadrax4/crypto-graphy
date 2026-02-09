package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	msg := []byte("there is nothing")

	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(msg, pubKey))

	// test with invalid msg
	assert.False(t, sig.Verify([]byte("nothing"), pubKey))

	// test with invalid pubkey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(msg, invalidPubKey))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()

	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
}

func TestNewPrivateKeyFromSeed(t *testing.T) {
	var (
		seed       = "5ca645d0c452e0f4c8715c8c0e9b08c45f84d1156aef9ef4c2875c064b4552ff"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "f704cba9b4835248211c1688eddb8c242af97cc0"
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))

	address := privKey.Public().Address()

	assert.Equal(t, address.String(), addressStr)
}
