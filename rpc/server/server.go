package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Up(n int, reply *int) error {
	*reply = n + 1
	return nil
}

func main() {
	rpc.Register(NewServer())
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Listen fail ", err)
		return
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.LocalAddr, "connect")
		go rpc.ServeConn(conn)
	}
}
