package model

import "time"

type Transaction struct {
	ID            int64
	DestinationID string `db:"destination_id" json:"destination_id"`
	TrxID         string `db:"trx_id" json:"trx_id"`
	Destination   Account
	Amount        int64
	UniqueNumber  int64 `db:"unique_number" json:"unique_number"`
	Type          TransactionType
	Status        TransactionStatus
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type DisbursementRequest struct {
	Account      Account `json:"account"`
	Amount       int64
	UniqueNumber int64 `json:"unique_number"`
}

type CallbackRequest struct {
	TransactionID string            `json:"transaction_id"`
	Status        TransactionStatus `json:"status"`
}

type TransactionStatus string
type TransactionType string

const (
	StatusPending TransactionStatus = "pending"
	StatusSuccess TransactionStatus = "success"
	StatusFailed  TransactionStatus = "failed"
)

const (
	Disbursement TransactionType = "disbursement"
)
