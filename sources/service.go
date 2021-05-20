package sources

import (
	"fmt"
	"strconv"
	"time"

	"github.com/lunnik9/product-api/domain"
	pe "github.com/lunnik9/product-api/product_errors"
	"github.com/lunnik9/product-api/sources/merch_repo"
	"github.com/lunnik9/product-api/sources/product_repo"
	"github.com/lunnik9/product-api/sources/waybill_repo"
	satori "github.com/satori/go.uuid"
)

const (
	WAYBILL_DOC_NUMBER_LEN = 6
)

type service struct {
	mr merch_repo.MerchRepo
	pr product_repo.ProductRepo
	wr waybill_repo.WaybillRepo
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
	MDleleteProducts(req *mDeleteProductsRequest) (*mDeleteProductsResponse, error)

	GetCategoryById(req *getCategoryByIdRequest) (*getCategoryByIdResponse, error)
	CreateCategory(req *createCategoryRequest) (*createCategoryResponse, error)
	UpdateCategory(req *updateCategoryRequest) (*updateCategoryResponse, error)
	DeleteCategory(req *deleteCategoryRequest) (*deleteCategoryResponse, error)
	FilterCategories(req *filterCategoriesRequest) (*filterCategoriesResponse, error)

	CreateWaybill(req *createWaybillRequest) (*createWaybillResponse, error)
	ConductWaybill(req *conductWaybillRequest) (*conductWaybillResponse, error)
	RollbackWaybill(req *rollbackWaybillRequest) (*rollbackWaybillResponse, error)
	DeleteWaybill(req *deleteWaybillRequest) (*deleteWaybillResponse, error)
	CreateWaybillProduct(req *createWaybillProductRequest) (*createWaybillProductResponse, error)
	UpdateWaybillProduct(req *updateWaybillProductRequest) (*updateWaybillProductResponse, error)
	DeleteWaybillProduct(req *deleteWaybillProductRequest) (*deleteWaybillProductResponse, error)
	GetListOfWaybillProducts(req *getListOfWaybillProductsRequest) (*getListOfWaybillProductsResponse, error)
}

func NewService(mr merch_repo.MerchRepo, pr product_repo.ProductRepo, wr waybill_repo.WaybillRepo) Service {
	return &service{
		mr: mr,
		pr: pr,
		wr: wr,
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

func (s *service) GetCategoryById(req *getCategoryByIdRequest) (*getCategoryByIdResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	category, err := s.pr.GetCategory(req.Id)
	if err != nil {
		return nil, err
	}

	return &getCategoryByIdResponse{*category}, nil
}

func (s *service) CreateCategory(req *createCategoryRequest) (*createCategoryResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.Category.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.Category.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	if req.Category.Name == "" {
		return nil, pe.New(409, "name cannot be empty")
	}

	id, err := s.pr.CreateCategory(req.Category)
	if err != nil {
		return nil, err
	}

	return &createCategoryResponse{id}, nil
}

func (s *service) UpdateCategory(req *updateCategoryRequest) (*updateCategoryResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	if req.Category.MerchantId == "" {
		return nil, pe.New(409, "merchant id cannot be empty")
	}

	if req.Category.StockId == "" {
		return nil, pe.New(409, "stock id cannot be empty")
	}

	if req.Category.Name == "" {
		return nil, pe.New(409, "name cannot be empty")
	}

	category, err := s.pr.UpdateCategory(req.Category)
	if err != nil {
		return nil, err
	}

	return &updateCategoryResponse{*category}, nil
}

func (s *service) DeleteCategory(req *deleteCategoryRequest) (*deleteCategoryResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	err = s.pr.DeleteCategory(req.Id)
	if err != nil {
		return nil, err
	}

	return &deleteCategoryResponse{req.Id}, nil
}
func (s *service) FilterCategories(req *filterCategoriesRequest) (*filterCategoriesResponse, error) {
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

	categories, err := s.pr.FilterCategory(req.Limit, req.Offset, req.MerchantId, req.StockId)
	if err != nil {
		return nil, err
	}

	return &filterCategoriesResponse{categories}, nil
}

func (s *service) MDleleteProducts(req *mDeleteProductsRequest) (*mDeleteProductsResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	err = s.pr.MDelete(req.Ids)
	if err != nil {
		return nil, err
	}

	return &mDeleteProductsResponse{req.Ids}, nil
}

func (s *service) CreateWaybill(req *createWaybillRequest) (*createWaybillResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	req.Waybill.Status = "draft"

	waybills, err := s.wr.Filter(-1, 0, req.Waybill.Type, "", req.Waybill.MerchantId, req.Waybill.StockId)
	if err != nil {
		return nil, err
	}

	req.Waybill.Number = strconv.Itoa(len(waybills))

	for len(req.Waybill.Number) < WAYBILL_DOC_NUMBER_LEN {
		req.Waybill.Number = "0" + req.Waybill.Number
	}

	id, err := s.wr.Create(req.Waybill)
	if err != nil {
		return nil, err
	}

	return &createWaybillResponse{id}, nil
}

func (s *service) ConductWaybill(req *conductWaybillRequest) (*conductWaybillResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	oldWaybill, err := s.wr.Get(req.Waybill.Id)
	if err != nil {
		return nil, err
	}

	if oldWaybill.Status != "draft" {
		return nil, pe.New(409, fmt.Sprintf("cannot conduct waybill w/ status %v", oldWaybill.Status))
	}

	products, err := s.wr.GetList(-1, 0, req.Waybill.Id)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		transfer := domain.Transfer{
			ProductId:     product.ProductId,
			SellingPrice:  product.SellingPrice,
			PurchasePrice: product.PurchasePrice,
			Amount:        product.Amount,
			Reason:        "received",
			Source:        "inwaybill",
			SourceId:      strconv.FormatInt(req.Waybill.Id, 10),
		}

		err = s.pr.SaveTransfer(transfer)
		if err != nil {
			return nil, err
		}
	}

	req.Waybill.Status = "active"
	req.Waybill.UpdatedOn = time.Now().UTC()

	waybill, err := s.wr.Update(req.Waybill)
	if err != nil {
		return nil, err
	}

	return &conductWaybillResponse{*waybill}, nil
}

func (s *service) RollbackWaybill(req *rollbackWaybillRequest) (*rollbackWaybillResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	oldWaybill, err := s.wr.Get(req.Waybill.Id)
	if err != nil {
		return nil, err
	}

	if oldWaybill.Status != "active" {
		return nil, pe.New(409, fmt.Sprintf("cannot rollback waybill w/ status %v", oldWaybill.Status))
	}

	products, err := s.wr.GetList(-1, 0, req.Waybill.Id)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		transfer := domain.Transfer{
			ProductId:     product.ProductId,
			SellingPrice:  product.SellingPrice,
			PurchasePrice: product.PurchasePrice,
			Amount:        -product.Amount,
			Reason:        "pulled",
			Source:        "outwaybill",
			SourceId:      strconv.FormatInt(req.Waybill.Id, 10),
		}

		err = s.pr.SaveTransfer(transfer)
		if err != nil {
			return nil, err
		}
	}

	req.Waybill.Status = "active"
	req.Waybill.UpdatedOn = time.Now().UTC()

	waybill, err := s.wr.Update(req.Waybill)
	if err != nil {
		return nil, err
	}

	return &rollbackWaybillResponse{*waybill}, nil
}

func (s *service) DeleteWaybill(req *deleteWaybillRequest) (*deleteWaybillResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	err = s.wr.Delete(req.Id)
	if err != nil {
		return nil, err
	}

	return &deleteWaybillResponse{req.Id}, nil
}

func (s *service) CreateWaybillProduct(req *createWaybillProductRequest) (*createWaybillProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	waybill, err := s.wr.Get(req.Product.WaybillId)
	if err != nil {
		return nil, err
	}

	if waybill.Status != "draft" {
		return nil, pe.New(409, fmt.Sprintf("cannot create product in %v waybill", waybill.Status))
	}

	_, err = s.wr.CreateProduct(req.Product)
	if err != nil {
		return nil, err
	}

	waybill.TotalCost += req.Product.PurchasePrice * req.Product.Amount
	waybill.UpdatedOn = time.Now().UTC()

	_, err = s.wr.Update(*waybill)
	if err != nil {
		return nil, err
	}

	return &createWaybillProductResponse{waybill.TotalCost}, nil
}
func (s *service) UpdateWaybillProduct(req *updateWaybillProductRequest) (*updateWaybillProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	waybill, err := s.wr.Get(req.Product.WaybillId)
	if err != nil {
		return nil, err
	}

	if waybill.Status != "draft" {
		return nil, pe.New(409, fmt.Sprintf("cannot update product in %v waybill", waybill.Status))
	}

	oldProduct, err := s.wr.GetProduct(req.Product.Id)
	if err != nil {
		return nil, err
	}

	_, err = s.wr.UpdateProduct(req.Product)
	if err != nil {
		return nil, err
	}

	waybill.TotalCost += (req.Product.PurchasePrice * req.Product.Amount) - (oldProduct.PurchasePrice * oldProduct.Amount)
	waybill.UpdatedOn = time.Now().UTC()

	_, err = s.wr.Update(*waybill)
	if err != nil {
		return nil, err
	}

	return &updateWaybillProductResponse{waybill.TotalCost}, nil
}

func (s *service) DeleteWaybillProduct(req *deleteWaybillProductRequest) (*deleteWaybillProductResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	product, err := s.wr.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}

	waybill, err := s.wr.Get(product.WaybillId)
	if err != nil {
		return nil, err
	}

	if waybill.Status != "draft" {
		return nil, pe.New(409, fmt.Sprintf("cannot update product in %v waybill", waybill.Status))
	}

	err = s.wr.DeleteProduct(req.Id)
	if err != nil {
		return nil, err
	}

	waybill.TotalCost -= product.PurchasePrice * product.Amount
	waybill.UpdatedOn = time.Now().UTC()

	_, err = s.wr.Update(*waybill)
	if err != nil {
		return nil, err
	}

	return &deleteWaybillProductResponse{waybill.TotalCost}, nil
}

func (s *service) GetListOfWaybillProducts(req *getListOfWaybillProductsRequest) (*getListOfWaybillProductsResponse, error) {
	err := s.mr.CheckRights(req.Authorization)
	if err != nil {
		return nil, err
	}

	products, err := s.wr.GetList(req.Limit, req.Offset, req.WaybillId)
	if err != nil {
		return nil, err
	}

	return &getListOfWaybillProductsResponse{products}, nil
}
