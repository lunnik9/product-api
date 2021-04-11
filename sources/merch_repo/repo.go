package merch_repo

import "github.com/lunnik9/product-api/domain"

type MerchRepo interface {
	GetMerchByNameAndPassword(mobile, password string) (*domain.Merchant, error)
	GetMerchByToken(token string) (*domain.Merchant, error)
	UpdateMerch(merch domain.Merchant) error
	CheckRights(token string) error
	CheckRightsWithMerch(merch domain.Merchant, token string) error
}
