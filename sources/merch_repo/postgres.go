package merch_repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
)

type MerchPostgres struct {
	db *pg.DB
}

func NewMerchPostgres(db *pg.DB) MerchPostgres {
	return MerchPostgres{db: db}
}

func (mr *MerchPostgres) GetMerchByNameAndPassword(mobile, password string) (*domain.Merchant, error) {
	var view domain.MerchantView

	err := mr.db.Model(&view).
		Table("merchant").
		Where("merchant.mobile = ? and merchant.password = ?", mobile, password).
		Select()
	if err != nil {
		return nil, err
	}

	merch := domain.MerchViewToDomain(view)

	return &merch, nil
}

func (mr *MerchPostgres) GetMerchByToken(token string) (*domain.Merchant, error) {
	var view domain.MerchantView

	err := mr.db.Model(&view).
		Table("merchant").
		Where("merchant.token = ?", token).
		Select()
	if err != nil {
		return nil, err
	}

	merch := domain.MerchViewToDomain(view)

	return &merch, nil
}

func (mr *MerchPostgres) UpdateMerch(merch domain.Merchant) error {
	view := domain.MerchDomainToView(merch)

	_, err := mr.db.Model(&view).WherePK().Update()
	if err != nil {
		return err
	}

	return nil
}

func (mr *MerchPostgres) CheckRights(token string) (bool, error) {
	_, err := mr.GetMerchByToken(token)
	if err != nil {
		return false, err
	}

	return true, nil
}
