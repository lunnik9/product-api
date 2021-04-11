package merch_repo

import "github.com/product-api/domain"

type MerchRepo interface {
	GetMerchByNameAndPassword(mobile, password string) (*domain.Merchant, error)
	GetMerchByToken(token string) (*domain.Merchant, error)
	UpdateMerch(merch domain.Merchant) error
	CheckRights(token string) (bool, error)
}
