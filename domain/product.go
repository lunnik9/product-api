package domain

import "time"

type ProductView struct {
	Barcode       string    `json:"barcode" pg:"barcode"`
	Name          string    `json:"name" pg:"name"`
	StockId       string    `json:"stock_id" pg:"stock_id"`
	Amount        float64   `json:"amount" pg:"amount"`
	PurchasePrice float64   `json:"purchase_price" pg:"purchase_price"`
	SellingPrice  float64   `json:"selling_price" pg:"selling_price"`
	CreatedOn     time.Time `json:"created_on" pg:"created_on"`
	UpdatedOn     time.Time `json:"updated_on" pg:"updated_on"`
}
