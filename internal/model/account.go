package model

type Account struct {
	ID         string `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Number     string `db:"number" json:"account_number"`
	OriginBank string `db:"origin_bank" json:"origin,omitempty"`
}

type CheckAccountRequest struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
}
