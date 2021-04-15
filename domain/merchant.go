package domain

import "time"

type MerchantView struct {
	tableName struct{} `pg:"merchant"`

	MerchantId   string    `pg:"merchant_id,pk"`
	MerchantName string    `pg:"merchant_name"`
	Password     string    `pg:"password"`
	Token        string    `pg:"token"`
	UpdateTime   time.Time `pg:"update_time"`
	TokenTTL     int       `pg:"token_ttl"`
	LastCheck    time.Time `pg:"last_check"`
	Mobile       string    `pg:"mobile"`
}

type Merchant struct {
	MerchantId   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	Password     string    `json:"password"`
	Token        string    `json:"token"`
	UpdateTime   time.Time `json:"update_time"`
	TokenTTL     int       `json:"token_ttl"`
	LastCheck    time.Time `json:"last_check"`
	Mobile       string    `json:"mobile"`
}

func MerchViewToDomain(view MerchantView) Merchant {
	return Merchant{
		MerchantId:   view.MerchantId,
		MerchantName: view.MerchantName,
		Password:     view.Password,
		Token:        view.Token,
		UpdateTime:   view.UpdateTime,
		TokenTTL:     view.TokenTTL,
		LastCheck:    view.LastCheck,
		Mobile:       view.Mobile,
	}
}
func MerchDomainToView(merchant Merchant) MerchantView {
	return MerchantView{
		MerchantId:   merchant.MerchantId,
		MerchantName: merchant.MerchantName,
		Password:     merchant.Password,
		Token:        merchant.Token,
		UpdateTime:   merchant.UpdateTime,
		TokenTTL:     merchant.TokenTTL,
		LastCheck:    merchant.LastCheck,
		Mobile:       merchant.Mobile,
	}
}

type Stock struct {
	StockId    string `json:"stock_id"`
	MerchantId string `json:"merchant_id"`
	StockName  string `json:"stock_name"`
}

type StockView struct {
	tableName struct{} `pg:"stock"`

	StockId    string `pg:"stock_id,pk"`
	MerchantId string `pg:"merchant_id"`
	StockName  string `pg:"stock_name"`
}

func StockViewToDomain(view StockView) Stock {
	return Stock{
		StockId:    view.StockId,
		MerchantId: view.MerchantId,
		StockName:  view.StockName,
	}
}

func StockDomainToView(stock Stock) StockView {
	return StockView{
		StockId:    stock.StockId,
		MerchantId: stock.MerchantId,
		StockName:  stock.StockName,
	}
}
