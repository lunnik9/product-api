package waybill_repo

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
)

type WaybillPostgres struct {
	db *pg.DB
}

func NewWaybillPostgres(db *pg.DB) WaybillPostgres {
	return WaybillPostgres{db: db}
}

func (wr *WaybillPostgres) Get(id int64) (*domain.Waybill, error) {
	var view = domain.WaybillView{Id: id}

	err := wr.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	waybill := domain.WaybillViewToDomain(view)

	return &waybill, nil
}

func (wr *WaybillPostgres) Delete(id int64) error {
	var view = domain.WaybillView{Id: id}

	_, err := wr.db.Model(&view).WherePK().Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (wr *WaybillPostgres) Create(waybill domain.Waybill) (int64, error) {
	waybill.CreatedOn = time.Now().UTC()
	waybill.UpdatedOn = time.Now().UTC()

	view := domain.WaybillDomainToView(waybill)

	_, err := wr.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	return view.Id, nil
}

func (wr *WaybillPostgres) Update(waybill domain.Waybill) (*domain.Waybill, error) {
	waybill.UpdatedOn = time.Now().UTC()

	view := domain.WaybillDomainToView(waybill)

	_, err := wr.db.Model(&view).WherePK().Update()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	return &waybill, nil
}

func (wr *WaybillPostgres) Filter(limit, offset int, docType, docNumber, merchantId, stockId string) ([]domain.Waybill, error) {
	var (
		views []domain.WaybillView
		resp  []domain.Waybill
		query string
	)

	query = "select * from waybill where merchant_id = ? and stock_id = ? and type = ?"

	if docNumber != "" {
		query += fmt.Sprintf(" and number = '%v'", docNumber)
	}

	query = query + " order by updated_on desc offset ?"

	if limit != -1 {
		query += fmt.Sprintf(" limit %v", limit)
	}

	_, err := wr.db.Query(&views, query, merchantId, stockId, docType, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		resp = append(resp, domain.WaybillViewToDomain(view))
	}

	return resp, nil
}

func (wr *WaybillPostgres) CreateProduct(product domain.WaybillProduct) (int64, error) {
	product.CreatedOn = time.Now().UTC()
	product.UpdatedOn = time.Now().UTC()

	view := domain.WaybillProductDomainToView(product)

	_, err := wr.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	return view.Id, nil
}

func (wr *WaybillPostgres) UpdateProduct(product domain.WaybillProduct) (*domain.WaybillProduct, error) {
	product.UpdatedOn = time.Now().UTC()

	view := domain.WaybillProductDomainToView(product)

	_, err := wr.db.Model(&view).WherePK().Update()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	return &product, nil
}

func (wr *WaybillPostgres) DeleteProduct(id int64) error {
	var view = domain.WaybillProductView{Id: id}

	_, err := wr.db.Model(&view).WherePK().Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}

func (wr *WaybillPostgres) GetList(limit, offset int, waybillId int64) ([]domain.WaybillProduct, error) {
	var (
		views []domain.WaybillProductView
		resp  []domain.WaybillProduct
		query string
	)

	query = "select * from waybill_product where waybill_id = ? order by updated_on desc offset ?"

	if limit != -1 {
		query += fmt.Sprintf(" limit %v", limit)
	}

	_, err := wr.db.Query(&views, query, waybillId, offset)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, view := range views {
		resp = append(resp, domain.WaybillProductViewToDomain(view))
	}

	return resp, nil
}

func (wr *WaybillPostgres) GetProduct(id int64) (*domain.WaybillProduct, error) {
	var view = domain.WaybillProductView{Id: id}

	err := wr.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	product := domain.WaybillProductViewToDomain(view)

	return &product, nil
}
