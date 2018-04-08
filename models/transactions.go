package models

import "time"

// Transaction struct describe the movement of stock in or out of a store
// This could be movement of stock between stores or purchase of stock from supplier etc
type Transaction struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionItem struct {
	ID            string
	TransactionID string
	ItemID        string
	Qty           int
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type InternalTransaction struct {
	Transaction
	Details       string
	Requester     string
	Approver      string
	Status        string
	Type          string
	RequestDate   time.Time
	ApprovedDate  time.Time
	DeliveredDate time.Time
}

type InternalTransactionItem struct {
	TransactionItem
	Status  string
	Remarks string
}
