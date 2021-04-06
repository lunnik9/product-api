package merch_repo

import "github.com/product-api/domain"

type MerchRepo interface {
	GetMerchByNameAndPassword(mobile, password string) (domain.MerchantView, error)
	UpdateMerch(merch domain.MerchantView) error
}
