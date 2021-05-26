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
		views      []domain.OrderItemView
		query      = "select * from order_product where order_id= ?"
		productIds []int64
	)

	_, err := or.db.Query(&views, query, orderId)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, item := range views {
		productIds = append(productIds, item.ProductId)
	}

	if len(views) != 0 {
		productMap, err := or.mGetProductsMap(productIds)
		if err != nil {
			return nil, err
		}

		for i := range views {
			views[i].Barcode = productMap[views[i].ProductId].Barcode
			views[i].Name = productMap[views[i].ProductId].Name
		}

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

func (or *OrderPostgres) mGetProductsMap(ids []int64) (map[int64]domain.Product, error) {
	var (
		views []domain.ProductView
		resp  = make(map[int64]domain.Product)
	)

	err := or.db.Model(&views).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		resp[view.Id] = domain.ProductViewToDomain(view)
	}

	return resp, nil
}
