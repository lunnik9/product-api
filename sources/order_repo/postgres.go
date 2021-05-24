package order_repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
)

type OrderPostgres struct {
	db *pg.DB
}

func NewOrderPostgres(db *pg.DB) OrderPostgres {
	return OrderPostgres{db: db}
}

func (pr *OrderPostgres) Get(id int64) (*domain.Order, error) {
	//var (
	//	view domain.OrderView
	//	items []domain.OrderItemView
	//	order domain.Order
	//)

	return nil, nil
}
