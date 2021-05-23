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

	_, err := mr.db.Query(&view, "select * from merchant where mobile= ? and password = ?", mobile, password)
	//err := mr.db.Model(&view).
	//	Where("merchant.mobile = ?", mobile).
	//	Where("and merchant.password = ?", password).
	//	Select()
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	merch := domain.MerchViewToDomain(view)

	return &merch, nil
}

func (mr *MerchPostgres) GetMerchByToken(token string) (*domain.Merchant, error) {
	var view domain.MerchantView

	_, err := mr.db.Query(&view, "select * from merchant where token = ?", token)
	//err := mr.db.Model(&view).
	//	Where("merchant.token = ?", token).
	//	Select()
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

func (mr *MerchPostgres) GetStocksOfMerchant(merchId string) ([]domain.Stock, error) {
	var (
		stockViews []domain.StockView
		stocks     []domain.Stock
		//query string
	)

	//query=""
	err := mr.db.Model(&stockViews).Select()
	if err != nil {
		return nil, pe.New(503, err.Error())
	}

	for _, v := range stockViews {
		stocks = append(stocks, domain.StockViewToDomain(v))
	}

	return stocks, nil
}

func (mr *MerchPostgres) GetListOfCashBoxes(merchId, stockId string) ([]domain.CashBox, error) {
	var (
		cashBoxViews []domain.CashBoxView
		cashBoxes    []domain.CashBox
	)

	query := "select * from cash_box where merchant_id = ? and stock_id = ?"

	_, err := mr.db.Query(&cashBoxViews, query, merchId, stockId)
	if err != nil {
		return nil, pe.New(409, err.Error())
	}

	for _, v := range cashBoxViews {
		cashBoxes = append(cashBoxes, domain.CashBoxViewToDomain(v))
	}

	return cashBoxes, nil
}
