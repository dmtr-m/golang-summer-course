package main

import (
	"context"
	"flag"
	"fmt"
	"gRPCProject/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Command struct {
	Type    string
	Name    string
	NewName string
	Amount  int
}

func Execute(command Command) error {
	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	client := proto.NewBankAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch command.Type {
	case "create":
		return CreateAccountExecutable(command, client, ctx)

	case "delete":
		return DeleteAccountExecutable(command, client, ctx)

	case "change_name":
		return ChangeAccountNameExecutable(command, client, ctx)

	case "change_balance":
		return ChangeAccountBalanceExecutable(command, client, ctx)

	case "get":
		return GetAccountExecutable(command, client, ctx)

	default:
		return fmt.Errorf("unknown command %s", command.Type)
	}
}

func CreateAccountExecutable(command Command, client proto.BankAccountServiceClient, ctx context.Context) error {
	res, err := client.CreateAccount(ctx, &proto.CreateAccountRequest{Name: command.Name, Amount: int64(command.Amount)})
	if err != nil {
		return err
	}

	fmt.Println(res.Message)
	return nil
}

func DeleteAccountExecutable(command Command, client proto.BankAccountServiceClient, ctx context.Context) error {
	res, err := client.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: command.Name})
	if err != nil {
		return err
	}

	fmt.Println(res.Message)
	return nil
}

func ChangeAccountBalanceExecutable(command Command, client proto.BankAccountServiceClient, ctx context.Context) error {
	res, err := client.ChangeAccountBalance(ctx, &proto.ChangeAccountBalanceRequest{Name: command.Name, Amount: int64(command.Amount)})
	if err != nil {
		return err
	}

	fmt.Println(res.Message)
	return nil
}

func ChangeAccountNameExecutable(command Command, client proto.BankAccountServiceClient, ctx context.Context) error {
	res, err := client.ChangeAccountName(ctx, &proto.ChangeAccountNameRequest{Name: command.Name, NewName: command.NewName})
	if err != nil {
		return err
	}

	fmt.Println(res.Message)
	return nil
}

func GetAccountExecutable(command Command, client proto.BankAccountServiceClient, ctx context.Context) error {
	res, err := client.GetAccount(ctx, &proto.GetAccountRequest{Name: command.Name})
	if err != nil {
		return err
	}

	fmt.Println(res.Message)
	return nil
}

func main() {
	type_val := flag.String("cmd", "", "command to execute")
	name_val := flag.String("name", "", "name of account")
	new_name_val := flag.String("newname", "", "new name of account")
	amount_val := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	command := Command{
		Type:    *type_val,
		Name:    *name_val,
		NewName: *new_name_val,
		Amount:  *amount_val,
	}

	if err := Execute(command); err != nil {
		panic(err)
	}
}
