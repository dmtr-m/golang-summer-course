package main

import (
	"fmt"
	"gRPCProject/accounts"
	"gRPCProject/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	handler := accounts.New()

	proto.RegisterBankAccountServiceServer(server, handler)

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
