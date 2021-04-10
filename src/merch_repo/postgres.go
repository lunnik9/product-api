package merch_repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/product-api/domain"
)

type MerchPostgres struct {
	db *pg.DB
}

func (mr *MerchPostgres) GetMerchByNameAndPassword(mobile, password string) (domain.MerchantView, error) {
	var merch domain.MerchantView
	err := mr.db.Model(&merch).
		Table("merchant").
		Where("merchant.mobile = ? and merchant.password = ?", mobile, password).
		Select()
	if err != nil {
		panic(err)
	}

	return merch, nil
}

func (mr *MerchPostgres) GetMerchByToken(token string) (domain.MerchantView, error) {
	return domain.MerchantView{}, nil
}
func (mr *MerchPostgres) UpdateMerch(merch domain.MerchantView) error {
	return nil
}
func (mr *MerchPostgres) CheckRights(token string) (bool, error) {
	return false, nil
}
