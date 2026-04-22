package node

import (
	"context"
	"net"

	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
	"google.golang.org/grpc"
)

type Node struct {
	peers map[net.Addr]*grpc.ClientConn
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) HandleTransaction(
	ctx context.Context,
	tx *proto.Transaction,
) (*proto.None, error) {
	return nil, nil
}
s