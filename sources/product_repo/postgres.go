package product_repo

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
)

type ProductPostgres struct {
	db *pg.DB
}

func NewProductPostgres(db *pg.DB) ProductPostgres {
	return ProductPostgres{db: db}
}

func (pr *ProductPostgres) Get(id int64) (*domain.Product, error) {
	var view = domain.ProductView{Id: id}

	err := pr.db.Model(&view).WherePK().Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	product := domain.ProductViewToDomain(view)

	return &product, nil
}
func (pr *ProductPostgres) Delete(id int64) error {

	var view = domain.ProductView{Id: id}

	_, err := pr.db.Model(&view).WherePK().Delete()
	if err != nil {
		return pe.New(409, err.Error())
	}

	return nil
}
func (pr *ProductPostgres) Create(product domain.Product) (int64, error) {
	product.CreatedOn = time.Now().UTC()
	product.UpdatedOn = time.Now().UTC()

	view := domain.ProductDomainToView(product)

	_, err := pr.db.Model(&view).Returning("id").Insert()
	if err != nil {
		return 0, pe.New(409, err.Error())
	}

	return view.Id, nil
}
func (pr *ProductPostgres) Update(product domain.Product) (*domain.Product, error) {
	product.UpdatedOn = time.Now().UTC()

	view := domain.ProductDomainToView(product)

	_, err := pr.db.Model(&view).WherePK().Update()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	return &product, nil
}
func (pr *ProductPostgres) Filter(limit, offset int, merchantId, stockId, name, barcode string) ([]domain.Product, error) {
	return nil, nil
}
