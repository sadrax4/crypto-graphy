package types

import (
	"crypto/sha256"

	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
	pb "google.golang.org/protobuf/proto"
)

func HashTransaction(t *proto.Transaction) []byte {
	b, err := pb.Marshal(t)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
