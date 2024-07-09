package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountBalanceRequest struct {
	Name      string `json:"name"`
	NewAmount int    `json:"new_amount"`
}

type ChangeAccountNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}
