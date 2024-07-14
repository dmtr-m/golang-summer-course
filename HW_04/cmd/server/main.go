package main

import (
	"context"
	"fmt"
	"gRPCProject/accounts"
	"gRPCProject/proto"
	"net"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	db_conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/BankAccountServer")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	handler := accounts.New(db_conn)

	proto.RegisterBankAccountServiceServer(server, handler)

	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
