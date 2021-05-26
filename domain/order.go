package domain

import "time"

type Order struct {
	Id         int64       `json:"id"`
	Timestamp  time.Time   `json:"timestamp"`
	MerchantId string      `json:"merchant_id"`
	StockId    string      `json:"stock_id"`
	CashBoxId  string      `json:"cash_box_id"`
	OrderItems []OrderItem `json:"order_items,omitempty"`
	TotalSum   float64     `json:"total_sum"`
	PayType    string      `json:"pay_type"`
}

type OrderItem struct {
	ProductId     int64   `json:"product_id"`
	Amount        float64 `json:"amount"`
	SellingPrice  float64 `json:"selling_price"`
	PurchasePrice float64 `json:"purchase_price"`
	Barcode       string  `json:"barcode"`
}

type OrderItemView struct {
	tableName struct{} `pg:"order_product"`

	OrderId       int64   `pg:"order_id"`
	ProductId     int64   `pg:"product_id"`
	Amount        float64 `pg:"amount"`
	SellingPrice  float64 `pg:"selling_price"`
	PurchasePrice float64 `pg:"purchase_price"`
}

type OrderView struct {
	tableName struct{} `pg:"online_order"`

	Id         int64     `pg:"id"`
	Timestamp  time.Time `pg:"timestamp"`
	MerchantId string    `pg:"merchant_id"`
	StockId    string    `pg:"stock_id"`
	CashBoxId  string    `pg:"cash_box_id"`
	TotalSum   float64   `pg:"total_sum"`
	PayType    string    `pg:"pay_type"`
	barcode    string
}

func OrderViewToDomain(view OrderView, items []OrderItemView) Order {
	var order = Order{
		Id:         view.Id,
		Timestamp:  view.Timestamp,
		MerchantId: view.MerchantId,
		StockId:    view.StockId,
		CashBoxId:  view.CashBoxId,
		TotalSum:   view.TotalSum,
		PayType:    view.PayType,
	}
	for _, item := range items {
		order.OrderItems = append(order.OrderItems, OrderItem{
			ProductId:     item.ProductId,
			Amount:        item.Amount,
			SellingPrice:  item.SellingPrice,
			PurchasePrice: item.PurchasePrice,
		})
	}
	return order
}

func OrderDomainToView(order Order) (OrderView, []OrderItemView) {
	var items []OrderItemView

	for _, item := range order.OrderItems {
		items = append(items, OrderItemView{
			OrderId:       order.Id,
			ProductId:     item.ProductId,
			Amount:        item.Amount,
			SellingPrice:  item.SellingPrice,
			PurchasePrice: item.PurchasePrice,
		})
	}

	return OrderView{
		Id:         order.Id,
		Timestamp:  order.Timestamp,
		MerchantId: order.MerchantId,
		StockId:    order.StockId,
		CashBoxId:  order.CashBoxId,
		TotalSum:   order.TotalSum,
		PayType:    order.PayType,
	}, items
}
