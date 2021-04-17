package product_repo

import "github.com/lunnik9/product-api/domain"

type ProductRepo interface {
	Get(id int64) (*domain.Product, error)
	Delete(id int64) error
	Create(product domain.Product) (int64, error)
	Update(product domain.Product) (*domain.Product, error)
	Filter(limit, offset int, merchantId, stockId, name string) ([]domain.Product, error)
}
