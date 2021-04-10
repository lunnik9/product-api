package domain

import "time"

type MerchantView struct {
	tableName struct{} `pg:"merchant"`

	MerchantId   string    `json:"merchant_id" pg:""`
	MerchantName string    `json:"merchant_name" pg:"merchant_name"`
	Password     string    `json:"password" pg:"password"`
	Token        string    `json:"token" pg:"token"`
	UpdateTime   time.Time `json:"update_time" pg:"update_time"`
	TokenTTL     int       `json:"token_ttl" pg:"token_ttl"`
	LastCheck    time.Time `json:"last_check" pg:"last_check"`
	Mobile       string    `json:"mobile" pg:"mobile"`
}
