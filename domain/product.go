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
	CategoryId    int64     `pg:"category_id"`
	//CategoryView  *CategoryView `pg:"rel:has-one"`
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
	Category      string    `json:"category"`
	CategoryId    int64     `json:"category_id"`
}

func ProductViewToDomain(view ProductView) Product {
	var product = Product{
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
		CategoryId:    view.CategoryId,
	}
	//if view.CategoryView != nil {
	//	product.Category = view.CategoryView.Name
	//}

	return product
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
		CategoryId:    product.CategoryId,
	}
}

type Category struct {
	Id         int64  `json:"id"`
	MerchantId string `json:"merchant_id"`
	StockId    string `json:"stock_id"`
	Name       string `json:"name"`
}

type CategoryView struct {
	tableName struct{} `pg:"category"`

	Id         int64  `pg:"id"`
	MerchantId string `pg:"merchant_id"`
	StockId    string `pg:"stock_id"`
	Name       string `pg:"name"`
}

func CategoryViewToDomain(view CategoryView) Category {
	return Category{
		Id:         view.Id,
		MerchantId: view.MerchantId,
		StockId:    view.StockId,
		Name:       view.Name,
	}
}

func CategoryDomainToView(category Category) CategoryView {
	return CategoryView{
		Id:         category.Id,
		MerchantId: category.MerchantId,
		StockId:    category.StockId,
		Name:       category.Name,
	}
}

type Transfer struct {
	Id            int64     `json:"id"`
	ProductId     int64     `json:"product_id"`
	SellingPrice  float64   `json:"selling_price"`
	PurchasePrice float64   `json:"purchase_price"`
	Amount        float64   `json:"amount"`
	Reason        string    `json:"reason"`
	Source        string    `json:"source"`
	SourceId      string    `json:"source_id"`
	Timestamp     time.Time `pg:"timestamp"`
}

type TransferView struct {
	tableName struct{} `pg:"transfer"`

	Id            int64     `pg:"id,pk"`
	ProductId     int64     `pg:"product_id"`
	SellingPrice  float64   `pg:"selling_price"`
	PurchasePrice float64   `pg:"purchase_price"`
	Amount        float64   `pg:"amount"`
	Reason        string    `pg:"reason"`
	Source        string    `pg:"source"`
	SourceId      string    `pg:"source_id"`
	Timestamp     time.Time `pg:"timestamp"`
}

func TransferDomainToView(transfer Transfer) TransferView {
	return TransferView{
		Id:            transfer.Id,
		ProductId:     transfer.ProductId,
		SellingPrice:  transfer.SellingPrice,
		PurchasePrice: transfer.PurchasePrice,
		Amount:        transfer.Amount,
		Reason:        transfer.Reason,
		Source:        transfer.Source,
		SourceId:      transfer.SourceId,
		Timestamp:     transfer.Timestamp,
	}
}

func TransferViewToDomain(view TransferView) Transfer {
	return Transfer{
		Id:            view.Id,
		ProductId:     view.ProductId,
		SellingPrice:  view.SellingPrice,
		PurchasePrice: view.PurchasePrice,
		Amount:        view.Amount,
		Reason:        view.Reason,
		Source:        view.Source,
		SourceId:      view.SourceId,
		Timestamp:     view.Timestamp,
	}
}
