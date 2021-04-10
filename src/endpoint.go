package src

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type loginResponse struct {
	MerchantId string `json:"merchant_id"`
	Token      string `json:"token"`
}

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		resp, err := s.Login(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getRefreshTokenRequest struct {
	Token string `json:"token"`
}

type getRefreshTokenResponse struct {
	Token string `json:"token"`
}

func makeGetRefreshTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRefreshTokenRequest)
		resp, err := s.GetRefreshToken(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
