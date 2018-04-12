package models

import (
	"time"
)

// Item describe a single item metadata in the store
type Item struct {
	ID            string
	Name          string
	Description   string
	CategoryID    string
	SubCategoryID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Category describes a product category metadata in the store
type Category struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// SubCategory describes a product subcategory metadata in the store
type SubCategory struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
