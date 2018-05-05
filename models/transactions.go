package models

import "time"

// Receipt is a grouping of transactions
type Receipt struct {
	ID          string
	Transaction []Transaction
}

// Transaction is a a representation of movement of goods on the product
// They can represent various types of items - e.g. Actual items, Promotional codes, Discounts etc
type Transaction struct {
	ID        string
	Item      Item
	ItemID    string
	Qty       int
	Price     int
	Store     Store
	StoreID   string
	ReceiptID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate function for the transaction item struct. This does not involve manipulating the data
// of the transaction item.
// We would want to ensure the follow items:
// - Qty is not zero, it should be a 1 or -1 but never 0
// - Price can be zero
func (t Transaction) Validate() error { return nil }

type InternalTransaction struct {
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
	Status  string
	Remarks string
}
