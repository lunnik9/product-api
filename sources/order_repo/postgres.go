package order_repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
)

type OrderPostgres struct {
	db *pg.DB
}

func NewOrderPostgres(db *pg.DB) OrderPostgres {
	return OrderPostgres{db: db}
}

func (or *OrderPostgres) Get(id int64) (*domain.Order, error) {
	var (
		view  = domain.OrderView{Id: id}
		order domain.Order
	)

	err := or.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	items, err := or.getOrderItemViews(id)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	order = domain.OrderViewToDomain(view, items)

	return &order, nil
}

func (or *OrderPostgres) getOrderItemViews(orderId int64) ([]domain.OrderItemView, error) {
	var (
		views []domain.OrderItemView
		query = "select * from order_product where order_id= ?"
	)

	_, err := or.db.Query(&views, query, orderId)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	return views, nil
}

func (or *OrderPostgres) GetList(stockId, merchantId string, limit, offset int) ([]domain.Order, error) {
	var (
		views  []domain.OrderView
		orders []domain.Order
		query  = "select * from online_order where merchant_id =? and stock_id = ? limit ? offset ?"
	)

	_, err := or.db.Query(&views, query, merchantId, stockId, limit, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		orders = append(orders, domain.OrderViewToDomain(view, nil))
	}

	return orders, nil
}

func (or *OrderPostgres) Save(order domain.Order) (int64, error) {

	view, items := domain.OrderDomainToView(order)

	_, err := or.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	for _, item := range items {
		item.OrderId = view.Id
		err = or.saveItem(item)
		if err != nil {
			return 0, err
		}
	}

	return view.Id, nil
}

func (or *OrderPostgres) saveItem(view domain.OrderItemView) error {
	_, err := or.db.Model(&view).Insert()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}
