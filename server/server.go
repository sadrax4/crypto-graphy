package server

import "net"

type Server struct {
	listenAddr string
	ln         net.Listener
}

func New(listenAddr string) (*Server, error) {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	return &Server{
		listenAddr: listenAddr,
		ln:         ln,
	}, err
}
