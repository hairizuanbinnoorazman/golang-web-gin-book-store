package models

import "time"

// Transaction struct describe the movement of stock in or out of a store
// This could be movement of stock between stores or purchase of stock from supplier etc
type Transaction struct {
	ID        string
	StoreID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate function checks to ensure the Transaction struct is valid
func (t Transaction) Validate() error { return nil }

type TransactionItem struct {
	ID            string
	TransactionID string
	ItemID        string
	Qty           int
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (t TransactionItem) Validate() error { return nil }

type InternalTransaction struct {
	Transaction
	Details       string
	RequesterID   string
	ApproverID    string
	Status        string
	Type          string
	RequestDate   time.Time
	ApprovedDate  time.Time
	DeliveredDate time.Time
}

func (i InternalTransaction) Validate() error { return nil }

type InternalTransactionItem struct {
	TransactionItem
	Status  string
	Remarks string
}
