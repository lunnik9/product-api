package sources

import (
	"fmt"
	"time"

	pe "github.com/lunnik9/product-api/product_errors"
	"github.com/lunnik9/product-api/sources/merch_repo"
	"github.com/lunnik9/product-api/sources/product_repo"
	satori "github.com/satori/go.uuid"
)

type service struct {
	mr merch_repo.MerchRepo
	pr product_repo.ProductRepo
}

type Service interface {
	Login(req *loginRequest) (*loginResponse, error)
	GetRefreshToken(req *getRefreshTokenRequest) (*getRefreshTokenResponse, error)
	ListMerchantStocks(req *listMerchantStocksRequest) (*listMerchantStocksResponse, error)
	GetListOfCashBoxes(req *getListOfCashBoxesRequest) (*getListOfCashBoxesResponse, error)

	GetProductById(req *getProductByIdRequest) (*getProductByIdResponse, error)
	CreateProduct(req *createProductRequest) (*createProductResponse, error)
	UpdateProduct(req *updateProductRequest) (*updateProductResponse, error)
	DeleteProduct(req *deleteProductRequest) (*deleteProductResponse, error)
	FilterProducts(req *filterProductsRequest) (*filterProductsResponse, error)
}

func NewService(mr merch_repo.MerchRepo, pr product_repo.ProductRepo) Service {
	return &service{
		mr: mr,
		pr: pr,
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

func (s *service) GetProductById(req *getProductByIdRequest) (*getProductByIdResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	product, err := s.pr.Get(req.Id)
	if err != nil {
		return nil, err
	}

	return &getProductByIdResponse{*product}, nil
}

func (s *service) CreateProduct(req *createProductRequest) (*createProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.Product.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.Product.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	if req.Product.Name == "" {
		return nil, pe.New(409, "name cannot be empty")
	}

	if req.Product.Barcode == "" {
		return nil, pe.New(409, "barcode cannot be empty")
	}

	foundProduct, _ := s.pr.GetProductByBarcode(req.Product.MerchantId, req.Product.StockId, req.Product.Barcode)
	if foundProduct != nil {
		return nil, pe.New(409, fmt.Sprintf("product with barcode %v already exists", req.Product.Barcode))
	}

	id, err := s.pr.Create(req.Product)
	if err != nil {
		return nil, err
	}

	return &createProductResponse{id}, nil
}

func (s *service) UpdateProduct(req *updateProductRequest) (*updateProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.Product.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.Product.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	if req.Product.Name == "" {
		return nil, pe.New(409, "name cannot be empty")
	}

	if req.Product.Barcode == "" {
		return nil, pe.New(409, "barcode cannot be empty")
	}

	product, err := s.pr.Update(req.Product)
	if err != nil {
		return nil, err
	}

	return &updateProductResponse{*product}, nil
}

func (s *service) DeleteProduct(req *deleteProductRequest) (*deleteProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	err = s.pr.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &deleteProductResponse{req.Id}, nil
}

func (s *service) FilterProducts(req *filterProductsRequest) (*filterProductsResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	products, err := s.pr.Filter(req.Limit, req.Offset, req.MerchantId, req.StockId, req.Name)
	if err != nil {
		return nil, err
	}

	return &filterProductsResponse{products}, nil
}

func (s *service) GetListOfCashBoxes(req *getListOfCashBoxesRequest) (*getListOfCashBoxesResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	cashBoxes, err := s.mr.GetListOfCashBoxes(req.MerchantId, req.StockId)
	if err != nil {
		return nil, err
	}

	return &getListOfCashBoxesResponse{cashBoxes}, nil
}
