package waybill_repo

import (
	"github.com/lunnik9/product-api/domain"
)

type WaybillRepo interface {
	Get(id int64) (*domain.Waybill, error)
	Delete(id int64) error
	Create(waybill domain.Waybill) (int64, error)
	Update(waybill domain.Waybill) (*domain.Waybill, error)
	Filter(limit, offset int, docType, docNumber, merchantId, stockId string) ([]domain.Waybill, error)

	GetProduct(id int64) (*domain.WaybillProduct, error)
	CreateProduct(product domain.WaybillProduct) (int64, error)
	UpdateProduct(product domain.WaybillProduct) (*domain.WaybillProduct, error)
	DeleteProduct(id int64) error
	GetList(limit, offset int, waybillId int64) ([]domain.WaybillProduct, error)
}
