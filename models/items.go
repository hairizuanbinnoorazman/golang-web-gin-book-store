package models

import (
	"time"
)

// Item describe a single item metadata in the store
type Item struct {
	ID            string
	Name          string
	Description   string
	Category      Category
	CategoryID    string
	SubCategory   SubCategory
	SubCategoryID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (i Item) validateItemName() error {
	return nil
}

func (i Item) validateItemCategory() error {
	return nil
}

func (i Item) validateItemSubCategory() error {
	return nil
}

func (i Item) Validate() error {
	return nil
}

// Category describes a product category metadata in the store
type Category struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c Category) validateCategoryName() error {
	return nil
}

func (c Category) validateCategoryDescription() error {
	return nil
}

func (c Category) Validate() error {
	return nil
}

// SubCategory describes a product subcategory metadata in the store
type SubCategory struct {
	ID          string
	Category    Category
	CategoryID  string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s SubCategory) validateSubCategoryName() error {
	return nil
}

func (s SubCategory) validateSubCategoryDescription() error {
	return nil
}

func (s SubCategory) Validate() error {
	return nil
}

type ShoppingBasket struct {
	Items []Item
}

type Promotion interface {
	RunPromotion()
}

func (sb ShoppingBasket) ApplyPromotion(a Promotion) Item {
	return Item{}
}
