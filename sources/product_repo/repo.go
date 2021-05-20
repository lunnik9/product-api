package product_repo

import "github.com/lunnik9/product-api/domain"

type ProductRepo interface {
	Get(id int64) (*domain.Product, error)
	Delete(id int64) error
	Create(product domain.Product) (int64, error)
	Update(product domain.Product) (*domain.Product, error)
	Filter(limit, offset int, merchantId, stockId, name string) ([]domain.Product, error)
	GetProductByBarcode(merchantId, stockId, barcode string) (*domain.Product, error)
	MDelete(ids []int64) error

	GetCategory(id int64) (*domain.Category, error)
	DeleteCategory(id int64) error
	CreateCategory(category domain.Category) (int64, error)
	UpdateCategory(category domain.Category) (*domain.Category, error)
	FilterCategory(limit, offset int, merchantId, stockId string) ([]domain.Category, error)

	SaveTransfer(transfer domain.Transfer) error
	InsertTransfer(transfer domain.Transfer) error

	//SaveTransfer()
	//GetTransfers()
}
