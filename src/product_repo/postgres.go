package product_repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/product-api/domain"
)

type ProductPostgres struct {
	db *pg.DB
}

func (pr *ProductPostgres) Get(merchantId, barcode, stockId string) (*domain.Product, error) {
	return nil, nil

}
