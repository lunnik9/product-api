package sources

import (
	"time"

	"github.com/lunnik9/product-api/sources/merch_repo"
	satori "github.com/satori/go.uuid"
)

type service struct {
	mr merch_repo.MerchRepo
}

type Service interface {
	Login(req *loginRequest) (*loginResponse, error)
	GetRefreshToken(req *getRefreshTokenRequest) (*getRefreshTokenResponse, error)
	ListMerchantStocks(req *listMerchantStocksRequest) (*listMerchantStocksResponse, error)
}

func NewService(mr merch_repo.MerchRepo) Service {
	return &service{
		mr: mr,
	}
}

func (s *service) Login(req *loginRequest) (*loginResponse, error) {
	merch, err := s.mr.GetMerchByNameAndPassword(req.Mobile, req.Password)
	if err != nil {
		return nil, err
	}

	merch.Token = satori.NewV1().String()
	merch.LastCheck = time.Now().UTC()

	err = s.mr.UpdateMerch(*merch)
	if err != nil {
		return nil, err
	}

	return &loginResponse{merch.MerchantId, merch.MerchantName, merch.Token}, nil
}

func (s *service) GetRefreshToken(req *getRefreshTokenRequest) (*getRefreshTokenResponse, error) {
	merch, err := s.mr.GetMerchByToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = s.mr.CheckRightsWithMerch(*merch, req.Token)
	if err != nil {
		return nil, err
	}

	merch.Token = satori.NewV1().String()
	merch.LastCheck = time.Now().UTC()

	err = s.mr.UpdateMerch(*merch)
	if err != nil {
		return nil, err
	}

	return &getRefreshTokenResponse{merch.Token}, nil
}

func (s *service) ListMerchantStocks(req *listMerchantStocksRequest) (*listMerchantStocksResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	stocks, err := s.mr.GetStocksOfMerchant(req.MerchantId)
	if err != nil {
		return nil, err
	}

	return &listMerchantStocksResponse{stocks}, nil
}
