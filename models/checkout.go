package models

import "time"

type Items struct {
	ID       int64   `json:"id" db:"id"`
	SKU      string  `json:"sku" db:"sku"`
	Name     string  `json:"name" db:"name"`
	Price    float64 `json:"price" db:"price"`
	Currency string  `json:"currency" db:"currency"`
	Qty      int     `json:"qty" db:"qty"`
}

type PromotionItems struct {
	ID             int       `json:"id" db:"id"`
	ItemID         int       `json:"item_id" db:"item_id"`
	ItemSKU        string    `json:"item_sku" db:"item_sku"`
	MinQty         int       `json:"min_qty" db:"min_qty"`
	FreeItem       *string   `json:"free_item" db:"free_item"`
	Discount       *float64  `json:"discount" db:"discount"`
	IsCashback     bool      `json:"is_cashback" db:"is_cashback"`
	DirectCashback *string   `json:"direct_cashback" db:"direct_cashback"`
	MaxCount       *int      `json:"max_count" db:"max_count"`
	Detail         *string   `json:"detail" db:"detail"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date" db:"end_date"`
}

type Orders struct {
	ID        int     `json:"id" db:"id"`
	UserID    int     `json:"user_id" db:"user_id"`
	ItemID    int     `json:"item_id" db:"item_id"`
	ItemSku   string  `json:"item_sku" db:"item_sku"`
	ItemName  string  `json:"item_name" db:"item_name"`
	ItemPrice float64 `json:"item_price" db:"item_price"`
	ItemQty   int     `json:"item_qty" db:"item_qty"`
	Discount  float64 `json:"discount" db:"discount"`
	Total     float64 `json:"total" db:"total"`
}
