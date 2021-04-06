package src

import (
	"github.com/product-api/src/merch_repo"
	satori "github.com/satori/go.uuid"
)

type service struct {
	mr merch_repo.MerchRepo
}

type Service interface {
	Login(req *loginRequest) (*loginResponse, error)
}

func NewService() Service {
	return &service{}
}

func (s *service) Login(req *loginRequest) (*loginResponse, error) {
	merch, err := s.mr.GetMerchByNameAndPassword(req.Mobile, req.Password)
	if err != nil {
		return nil, err
	}

	merch.Token = satori.NewV1().String()

	err = s.mr.UpdateMerch(merch)
	if err != nil {
		return nil, err
	}

	return &loginResponse{merch.MerchantId, merch.Token}, nil
}
