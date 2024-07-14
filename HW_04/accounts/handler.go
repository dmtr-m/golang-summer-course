package accounts

import (
	"context"
	"fmt"
	"gRPCProject/proto"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func New(db_connection *pgx.Conn) *Handler {
	return &Handler{
		database: db_connection,
	}
}

type Handler struct {
	database *pgx.Conn
	proto.BankAccountServiceServer
}

func (h *Handler) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	log.Printf("Requested to create an account for %s with %d;", request.Name, request.Amount)

	if len(request.Name) == 0 {
		log.Printf("Rejected: invalid name;")
		return &proto.CreateAccountResponse{
			Message: "Invalid name",
		}, nil
	}

	if request.Amount < 0 {
		log.Printf("Rejected: negative amount;")
		return &proto.CreateAccountResponse{
			Message: "Amount can only be non-negative",
		}, nil
	}
	_, err := h.database.Exec(context.Background(), "INSERT INTO BankAccounts (name, balance) VALUES ($1, $2)", request.Name, request.Amount)

	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && "23505" == pgErr.Code {
			log.Printf("Rejected: account already exists;")
			return &proto.CreateAccountResponse{
				Message: "Account with such name already exists",
			}, nil
		}

		log.Printf("Rejected: database query failed;")
		return &proto.CreateAccountResponse{
			Message: "Database query failed",
		}, nil
	}

	log.Printf("Approved: account created;")

	return &proto.CreateAccountResponse{
		Message: "An account for " + request.Name + " was created. Current amount is " + fmt.Sprintf("%v", request.Amount),
	}, nil
}

func (h *Handler) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	log.Printf("Requested to create an account \"%s\";", request.Name)

	if len(request.Name) == 0 {
		log.Printf("Rejected: invalid name;")
		return &proto.DeleteAccountResponse{
			Message: "Invalid name",
		}, nil
	}

	req, err := h.database.Exec(context.Background(), "DELETE FROM BankAccounts WHERE name = $1", request.Name)

	if err != nil {
		log.Printf("Rejected: database query failed;")
		return &proto.DeleteAccountResponse{
			Message: "Database query failed",
		}, nil
	}

	if req.RowsAffected() == 0 {
		log.Printf("Rejected: account does not exist;")
		return &proto.DeleteAccountResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	log.Printf("Approved: account deleted;")

	return &proto.DeleteAccountResponse{
		Message: "Account \"" + request.Name + "\" was deleted.",
	}, nil
}

func (h *Handler) ChangeAccountBalance(ctx context.Context, request *proto.ChangeAccountBalanceRequest) (*proto.ChangeAccountBalanceResponse, error) {
	log.Printf("Requested to change an account %s balance, to %d;", request.Name, request.Amount)

	if len(request.Name) == 0 {
		log.Printf("Rejected: invalid name;")
		return &proto.ChangeAccountBalanceResponse{
			Message: "Invalid name",
		}, nil
	}

	if request.Amount < 0 {
		log.Printf("Rejected: negative amount;")
		return &proto.ChangeAccountBalanceResponse{
			Message: "Amount can only be non-negative",
		}, nil
	}

	req, err := h.database.Exec(context.Background(), "UPDATE BankAccounts SET balance = $1 WHERE name = $2", request.Amount, request.Name)

	if err != nil {
		log.Printf("Rejected: database query failed;")
		return &proto.ChangeAccountBalanceResponse{
			Message: "Database query failed",
		}, nil
	}

	if req.RowsAffected() == 0 {
		log.Printf("Rejected: account does not exist;")
		return &proto.ChangeAccountBalanceResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	log.Printf("Approved: account balance changed;")

	return &proto.ChangeAccountBalanceResponse{
		Message: "Account \"" + request.Name + "\" now has a balance of " + fmt.Sprintf("%v", request.Amount),
	}, nil
}

func (h *Handler) ChangeAccountName(ctx context.Context, request *proto.ChangeAccountNameRequest) (*proto.ChangeAccountNameResponse, error) {
	log.Printf("Requested to change an account name of %s, to %s;", request.Name, request.NewName)

	if len(request.Name) == 0 {
		log.Printf("Rejected: invalid name;")
		return &proto.ChangeAccountNameResponse{
			Message: "Invalid name",
		}, nil
	}

	if len(request.NewName) == 0 {
		log.Printf("Rejected: invalid new name;")
		return &proto.ChangeAccountNameResponse{
			Message: "Invalid new name",
		}, nil
	}

	req, err := h.database.Exec(context.Background(), "UPDATE BankAccounts SET name = $1 WHERE name = $2", request.NewName, request.Name)

	if err != nil {
		log.Printf("Rejected: database query failed;")
		return &proto.ChangeAccountNameResponse{
			Message: "Database query failed",
		}, nil
	}

	if req.RowsAffected() == 0 {
		log.Printf("Rejected: account does not exist;")
		return &proto.ChangeAccountNameResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	log.Printf("Approved: account renamed;")

	return &proto.ChangeAccountNameResponse{
		Message: "Account \"" + request.Name + "\" was renamed to \"" + request.NewName + "\"",
	}, nil
}

func (h *Handler) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	log.Printf("Requested to share an account %s;", request.Name)

	if len(request.Name) == 0 {
		log.Printf("Rejected: invalid name;")
		return &proto.GetAccountResponse{
			Message: "Invalid name",
		}, nil
	}

	var balance int
	req, err := h.database.Query(context.Background(), "SELECT balance FROM BankAccounts WHERE name = $1", request.Name)

	if err != nil {
		log.Printf("Rejected: database query failed;")
		return &proto.GetAccountResponse{
			Message: "Database query failed",
		}, nil
	}

	defer req.Close()

	if !req.Next() {
		log.Printf("Rejected: account does not exist;")
		return &proto.GetAccountResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	err = req.Scan(&balance)

	if err != nil {
		log.Printf("Rejected: database query failed;")
		return &proto.GetAccountResponse{
			Message: "Database query failed",
		}, nil
	}

	message := fmt.Sprintf("The account \"%s\" has a balance of %v", request.Name, balance)

	log.Printf("Approved: account shared;")

	return &proto.GetAccountResponse{
		Message: message,
	}, nil
}
