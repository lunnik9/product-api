package order_repo

import "github.com/lunnik9/product-api/domain"

type OrderRepo interface {
	Get(id int64) (*domain.Order, error)
	GetList(stockId, merchantId string, limit, offset int) ([]domain.Order, error)
	Save(order domain.Order) (int64, error)
}
