package domain

type ProductView struct {
	Barcode       string
	Name          string
	StockId       string
	Amount        float64
	PurchasePrice float64
	SellingPrice  float64
}
