package accounts

import (
	"httpProject/accounts/dto"
	"httpProject/accounts/models"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
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
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid Commamd")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Invalid name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "This username is already taken")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid command")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Invalid name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "Account with such name does not exist")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) ChangeAccountBalance(c echo.Context) error {
	var request dto.ChangeAccountBalanceRequest

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid command")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Invalid name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "Account with such name does not exist")
	}

	h.accounts[request.Name].Amount = request.NewAmount

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) ChangeAccountName(c echo.Context) error {
	var request dto.ChangeAccountNameRequest

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid command")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Invalid name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "Account with such name does not exist")
	}

	if _, ok := h.accounts[request.NewName]; ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "This username is already taken")
	}

	user := h.accounts[request.Name]
	user.Name = request.NewName
	delete(h.accounts, request.Name)
	h.accounts[request.NewName] = user

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccountDetails(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "Account with such name does not exist")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}
