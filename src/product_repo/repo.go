package product_repo

import "github.com/lunnik9/product-api/domain"

type ProductRepo interface {
	Get(merchantId, barcode, stockId string) (domain.ProductView, error)
	Delete(merchantId, barcode, stockId string) error
	Create(product domain.ProductView) error
	Update(product domain.ProductView) error
	Filter(limit, offset int, merchantId, stockId string)
}
