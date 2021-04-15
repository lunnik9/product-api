package domain

import "time"

type ProductView struct {
	tableName struct{} `pg:"product"`

	Id            int64     `pg:"id,pk"`
	Barcode       string    `pg:"barcode"`
	Name          string    `pg:"name"`
	StockId       string    `pg:"stock_id"`
	Amount        float64   `pg:"amount"`
	Unit          string    `pg:"unit"`
	MerchantId    string    `pg:"merchant_id"`
	PurchasePrice float64   `pg:"purchase_price"`
	SellingPrice  float64   `pg:"selling_price"`
	CreatedOn     time.Time `pg:"created_on"`
	UpdatedOn     time.Time `pg:"updated_on"`
}

type Product struct {
	Id            int64     `json:"id"`
	Barcode       string    `json:"barcode"`
	Name          string    `json:"name"`
	StockId       string    `json:"stock_id"`
	Amount        float64   `json:"amount"`
	Unit          string    `json:"unit"`
	MerchantId    string    `json:"merchant_id"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
}

func ProductViewToDomain(view ProductView) Product {
	return Product{
		Id:            view.Id,
		Barcode:       view.Barcode,
		Name:          view.Name,
		StockId:       view.StockId,
		Amount:        view.Amount,
		PurchasePrice: view.PurchasePrice,
		Unit:          view.Unit,
		MerchantId:    view.MerchantId,
		SellingPrice:  view.SellingPrice,
		CreatedOn:     view.CreatedOn,
		UpdatedOn:     view.UpdatedOn,
	}
}

func ProductDomainToView(product Product) ProductView {
	return ProductView{
		Id:            product.Id,
		Barcode:       product.Barcode,
		Name:          product.Name,
		StockId:       product.StockId,
		Amount:        product.Amount,
		Unit:          product.Unit,
		MerchantId:    product.MerchantId,
		PurchasePrice: product.PurchasePrice,
		SellingPrice:  product.SellingPrice,
		CreatedOn:     product.CreatedOn,
		UpdatedOn:     product.UpdatedOn,
	}
}
