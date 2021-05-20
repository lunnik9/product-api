package domain

import "time"

type Waybill struct {
	Id           int64      `json:"id"`
	MerchantId   string     `json:"merchant_id"`
	StockId      string     `json:"stock_id"`
	CreatedOn    time.Time  `json:"created_on"`
	UpdatedOn    time.Time  `json:"updated_on"`
	ReservedTime *time.Time `json:"reserved_time"`
	TotalCost    float64    `json:"total_cost"`
	Status       string     `json:"status"`
	Type         string     `json:"type"`
	Number       string     `json:"number"`
}

type WaybillView struct {
	tableName struct{} `pg:"waybill"`

	Id           int64      `pg:"id"`
	MerchantId   string     `pg:"merchant_id"`
	StockId      string     `pg:"stock_id"`
	CreatedOn    time.Time  `pg:"created_on"`
	UpdatedOn    time.Time  `pg:"updated_on"`
	ReservedTime *time.Time `pg:"reserved_time"`
	TotalCost    float64    `pg:"total_cost"`
	Status       string     `pg:"status"`
	Type         string     `pg:"type"`
	Number       string     `pg:"number"`
}

func WaybillViewToDomain(view WaybillView) Waybill {
	return Waybill{
		Id:           view.Id,
		MerchantId:   view.MerchantId,
		StockId:      view.StockId,
		CreatedOn:    view.CreatedOn,
		UpdatedOn:    view.UpdatedOn,
		ReservedTime: view.ReservedTime,
		TotalCost:    view.TotalCost,
		Status:       view.Status,
		Type:         view.Type,
		Number:       view.Number,
	}
}

func WaybillDomainToView(waybill Waybill) WaybillView {
	return WaybillView{
		Id:           waybill.Id,
		MerchantId:   waybill.MerchantId,
		StockId:      waybill.StockId,
		CreatedOn:    waybill.CreatedOn,
		UpdatedOn:    waybill.UpdatedOn,
		ReservedTime: waybill.ReservedTime,
		TotalCost:    waybill.TotalCost,
		Status:       waybill.Status,
		Type:         waybill.Type,
		Number:       waybill.Number,
	}
}

type WaybillProduct struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Barcode       string    `json:"barcode"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
	Amount        float64   `json:"amount"`
	WaybillId     int64     `json:"waybill_id"`
	ProductId     int64     `json:"product_id"`
}

type WaybillProductView struct {
	tableName struct{} `pg:"waybill_product"`

	Id            int64     `pg:"id"`
	Name          string    `pg:"name"`
	Barcode       string    `pg:"barcode"`
	PurchasePrice float64   `pg:"purchase_price"`
	SellingPrice  float64   `pg:"selling_price"`
	CreatedOn     time.Time `pg:"created_on"`
	UpdatedOn     time.Time `pg:"updated_on"`
	Amount        float64   `pg:"amount"`
	WaybillId     int64     `pg:"waybill_id"`
	ProductId     int64     `pg:"product_id"`
}

func WaybillProductDomainToView(product WaybillProduct) WaybillProductView {
	return WaybillProductView{
		Id:            product.Id,
		Name:          product.Name,
		Barcode:       product.Barcode,
		PurchasePrice: product.PurchasePrice,
		SellingPrice:  product.SellingPrice,
		CreatedOn:     product.CreatedOn,
		UpdatedOn:     product.UpdatedOn,
		Amount:        product.Amount,
		WaybillId:     product.WaybillId,
		ProductId:     product.ProductId,
	}
}

func WaybillProductViewToDomain(view WaybillProductView) WaybillProduct {
	return WaybillProduct{
		Id:            view.Id,
		Name:          view.Name,
		Barcode:       view.Barcode,
		PurchasePrice: view.PurchasePrice,
		SellingPrice:  view.SellingPrice,
		CreatedOn:     view.CreatedOn,
		UpdatedOn:     view.UpdatedOn,
		Amount:        view.Amount,
		WaybillId:     view.WaybillId,
		ProductId:     view.ProductId,
	}
}
