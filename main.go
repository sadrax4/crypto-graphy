package main

import (
	"log"
	"net"

	proto "github.com/sadrax4/crypto-graphy/github.com/sadrax4/crypto-graphy"
	node "github.com/sadrax4/crypto-graphy/node"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	opt := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opt...)
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterNodeServer(grpcServer, node)
	grpcServer.Serve(ln)
}

func makeTransaction() {
	kir
}
