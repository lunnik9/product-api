package sources

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lunnik9/product-api/domain"
)

type loginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type loginResponse struct {
	MerchantId   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	Token        string `json:"token"`
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

type listMerchantStocksRequest struct {
	Authorization string `json:"authorization"`
	MerchantId    string `json:"merchant_id"`
}

type listMerchantStocksResponse struct {
	Stocks []domain.Stock `json:"stocks"`
}

func makeListMerchantStocksEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listMerchantStocksRequest)
		resp, err := s.ListMerchantStocks(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getProductByIdRequest struct {
	Authorization string
	Id            int64
}

type getProductByIdResponse struct {
	Product domain.Product `json:"product"`
}

func makeGetProductByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIdRequest)
		resp, err := s.GetProductById(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type createProductRequest struct {
	Authorization string
	Product       domain.Product `json:"product"`
}

type createProductResponse struct {
	Id int64 `json:"id"`
}

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createProductRequest)
		resp, err := s.CreateProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type updateProductRequest struct {
	Authorization string
	Product       domain.Product `json:"product"`
}

type updateProductResponse struct {
	Product domain.Product `json:"product"`
}

func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		resp, err := s.UpdateProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type deleteProductRequest struct {
	Authorization string
	Id            int64
}

type deleteProductResponse struct {
	Id int64 `json:"id"`
}

func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		resp, err := s.DeleteProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
