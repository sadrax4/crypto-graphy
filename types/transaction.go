package types

import (
	"crypto/sha256"

	"github.com/sadrax4/crypto-graphy/crypto"
	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(
	pk *crypto.PrivateKey,
	tx *proto.Transaction,
) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(t *proto.Transaction) []byte {
	b, err := pb.Marshal(t)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
