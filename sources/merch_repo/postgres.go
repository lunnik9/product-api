package merch_repo

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
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
		return nil, pe.New(409, err.Error())
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
		return nil, pe.New(409, err.Error())
	}

	merch := domain.MerchViewToDomain(view)

	return &merch, nil
}

func (mr *MerchPostgres) UpdateMerch(merch domain.Merchant) error {
	view := domain.MerchDomainToView(merch)

	_, err := mr.db.Model(&view).WherePK().Update()
	if err != nil {
		return pe.New(503, err.Error())
	}

	return nil
}

func (mr *MerchPostgres) CheckRights(token string) error {
	merch, err := mr.GetMerchByToken(token)
	if err != nil {
		return pe.New(401, err.Error())
	}

	return mr.CheckRightsWithMerch(*merch, token)
}
func (mr *MerchPostgres) CheckRightsWithMerch(merch domain.Merchant, token string) error {
	timeout := merch.LastCheck.Add(time.Duration(merch.TokenTTL) * time.Second)

	if timeout.Before(time.Now().UTC()) {
		return pe.New(401, fmt.Sprintf("token %v timed out: %v", token, timeout))
	}
	return nil

}
