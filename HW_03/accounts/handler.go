package accounts

import (
	"context"
	"fmt"
	"gRPCProject/accounts/models"
	"gRPCProject/proto"
	"log"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
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

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()
		log.Printf("Rejected: username taken;")
		return &proto.CreateAccountResponse{
			Message: "This username is taken",
		}, nil
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

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

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		log.Printf("Rejected: account does not exist;")
		return &proto.DeleteAccountResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

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

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		log.Printf("Rejected: account does not exist;")
		return &proto.ChangeAccountBalanceResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	h.accounts[request.Name].Amount = request.Amount

	h.guard.Unlock()

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

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		log.Printf("Rejected: account does not exist;")
		return &proto.ChangeAccountNameResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	user := h.accounts[request.Name]
	user.Name = request.NewName
	delete(h.accounts, request.Name)
	h.accounts[request.NewName] = user

	h.guard.Unlock()

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

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		log.Printf("Rejected: account does not exist;")
		return &proto.GetAccountResponse{
			Message: "Account with such name does not exist",
		}, nil
	}

	message := "The account \"" + request.Name + "\" has a balance of " + fmt.Sprintf("%v", h.accounts[request.Name].Amount)

	h.guard.Unlock()

	log.Printf("Approved: account shared;")

	return &proto.GetAccountResponse{
		Message: message,
	}, nil
}
