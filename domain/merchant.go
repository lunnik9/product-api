package domain

import "time"

type MerchantView struct {
	MerchantId   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	Password     string    `json:"password"`
	Token        string    `json:"token"`
	UpdateTime   time.Time `json:"update_time"`
	TokenTTL     int       `json:"token_ttl"`
	LastCheck    time.Time `json:"last_check"`
	Mobile       string    `json:"mobile"`
}
