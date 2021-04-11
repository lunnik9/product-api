package merch_repo

import (
	"time"

	"github.com/lunnik9/product-api/domain"
)

type merchMock struct {
	Merchats []domain.MerchantView
}

func Init() merchMock {
	return merchMock{
		Merchats: []domain.MerchantView{
			{
				MerchantId:   "123",
				MerchantName: "arnur",
				Password:     "qwe",
				Token:        "",
				UpdateTime:   time.Now(),
				TokenTTL:     0,
				LastCheck:    time.Now(),
				Mobile:       "88005553535",
			},
			{
				MerchantId:   "456",
				MerchantName: "karina",
				Password:     "asd",
				Token:        "",
				UpdateTime:   time.Now(),
				TokenTTL:     0,
				LastCheck:    time.Now(),
				Mobile:       "228228228",
			},
			{
				MerchantId:   "789",
				MerchantName: "erasyl",
				Password:     "zxc",
				Token:        "",
				UpdateTime:   time.Now(),
				TokenTTL:     0,
				LastCheck:    time.Now(),
				Mobile:       "2128506",
			},
		},
	}
}

func (m merchMock) GetMerchByNameAndPassword(mobile, password string) (domain.MerchantView, error) {
	return domain.MerchantView{}, nil
}
func (m merchMock) UpdateMerch(merch domain.MerchantView) error {
	return nil
}
