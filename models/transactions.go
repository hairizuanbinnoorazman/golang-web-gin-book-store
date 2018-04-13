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
	requesterID   string
	approverID    string
	status        string
	Type          string
	requestDate   time.Time
	approvedDate  time.Time
	DeliveredDate time.Time
}

func (i InternalTransaction) Validate() error { return nil }

type InternalTransactionItem struct {
	TransactionItem
	Status  string
	Remarks string
}
