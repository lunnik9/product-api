package sources

import (
	"context"
	"time"

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

type filterProductsRequest struct {
	Authorization string
	Name          string `json:"name"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	MerchantId    string `json:"merchant_id"`
	StockId       string `json:"stock_id"`
}

type filterProductsResponse struct {
	Products []domain.Product `json:"products"`
}

func makeFilterProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(filterProductsRequest)
		resp, err := s.FilterProducts(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getListOfCashBoxesRequest struct {
	Authorization string
	MerchantId    string `json:"merchant_id"`
	StockId       string `json:"stock_id"`
}

type getListOfCashBoxesResponse struct {
	CashBoxes []domain.CashBox `json:"cash_boxes"`
}

func makeGetListOfCashBoxesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getListOfCashBoxesRequest)
		resp, err := s.GetListOfCashBoxes(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getCategoryByIdRequest struct {
	Authorization string
	Id            int64
}

type getCategoryByIdResponse struct {
	Category domain.Category `json:"category"`
}

func makeGetCategoryByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCategoryByIdRequest)
		resp, err := s.GetCategoryById(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type createCategoryRequest struct {
	Authorization string
	Category      domain.Category `json:"category"`
}

type createCategoryResponse struct {
	Id int64 `json:"id"`
}

func makeCreateCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createCategoryRequest)
		resp, err := s.CreateCategory(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type updateCategoryRequest struct {
	Authorization string
	Category      domain.Category `json:"category"`
}

type updateCategoryResponse struct {
	Category domain.Category `json:"category"`
}

func makeUpdateCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateCategoryRequest)
		resp, err := s.UpdateCategory(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type deleteCategoryRequest struct {
	Authorization string
	Id            int64
}

type deleteCategoryResponse struct {
	Id int64 `json:"id"`
}

func makeDeleteCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteCategoryRequest)
		resp, err := s.DeleteCategory(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type filterCategoriesRequest struct {
	Authorization string
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	MerchantId    string `json:"merchant_id"`
	StockId       string `json:"stock_id"`
}

type filterCategoriesResponse struct {
	Categories []domain.Category `json:"categories"`
}

func makeFilterCategoriesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(filterCategoriesRequest)
		resp, err := s.FilterCategories(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type mDeleteProductsRequest struct {
	Authorization string
	Ids           []int64 `json:"ids"`
}

type mDeleteProductsResponse struct {
	Ids []int64 `json:"ids"`
}

func makeMDeleteProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(mDeleteProductsRequest)
		resp, err := s.MDleteProducts(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type createWaybillRequest struct {
	Authorization string
	Waybill       domain.Waybill `json:"waybill"`
}

type createWaybillResponse struct {
	Id        int64     `json:"id"`
	Number    string    `json:"number"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

func makeCreateWaybillEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createWaybillRequest)
		resp, err := s.CreateWaybill(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type conductWaybillRequest struct {
	Authorization string
	Id            int64
}

type conductWaybillResponse struct {
	Waybill domain.Waybill `json:"waybill"`
}

func makeConductWaybillEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(conductWaybillRequest)
		resp, err := s.ConductWaybill(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type rollbackWaybillRequest struct {
	Authorization string
	Id            int64
}

type rollbackWaybillResponse struct {
	Waybill domain.Waybill `json:"waybill"`
}

func makeRollbackWaybillEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(rollbackWaybillRequest)
		resp, err := s.RollbackWaybill(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type deleteWaybillRequest struct {
	Authorization string
	Id            int64
}

type deleteWaybillResponse struct {
	Id int64 `json:"id"`
}

func makeDeleteWaybillEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteWaybillRequest)
		resp, err := s.DeleteWaybill(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type createWaybillProductRequest struct {
	Authorization string
	Product       domain.WaybillProduct `json:"product"`
}

type createWaybillProductResponse struct {
	TotalCost float64 `json:"total_cost"`
	Id        int64   `json:"id"`
}

func makeCreateWaybillProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createWaybillProductRequest)
		resp, err := s.CreateWaybillProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type updateWaybillProductRequest struct {
	Authorization string
	Product       domain.WaybillProduct `json:"product"`
}

type updateWaybillProductResponse struct {
	TotalCost float64 `json:"total_cost"`
}

func makeUpdateWaybillProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateWaybillProductRequest)
		resp, err := s.UpdateWaybillProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type deleteWaybillProductRequest struct {
	Authorization string
	Id            int64
}

type deleteWaybillProductResponse struct {
	TotalCost float64 `json:"total_cost"`
}

func makeDeleteWaybillProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteWaybillProductRequest)
		resp, err := s.DeleteWaybillProduct(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getListOfWaybillProductsRequest struct {
	Authorization string
	Limit         int   `json:"limit"`
	Offset        int   `json:"offset"`
	WaybillId     int64 `json:"waybill_id"`
}

type getListOfWaybillProductsResponse struct {
	Products []domain.WaybillProduct `json:"products"`
}

func makeGetListOfWaybillProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getListOfWaybillProductsRequest)
		resp, err := s.GetListOfWaybillProducts(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type filterWaybillsRequest struct {
	Authorization  string
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	WaybillType    string `json:"waybill_type"`
	DocumentNumber string `json:"document_number"`
	MerchantId     string `json:"merchant_id"`
	StockId        string `json:"stock_id"`
}

type filterWaybillsResponse struct {
	TotalHits int              `json:"total_hits"`
	Waybills  []domain.Waybill `json:"waybills"`
}

func makeFilterWaybillsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(filterWaybillsRequest)
		resp, err := s.FilterWaybills(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getWaybillByIdRequest struct {
	Authorization string
	Id            int64
}

type getWaybillByIdResponse struct {
	Waybill domain.Waybill `json:"waybill"`
}

func makeGetWaybillByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getWaybillByIdRequest)
		resp, err := s.GetWaybillById(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getWaybillProductByIdRequest struct {
	Authorization string
	Id            int64
}

type getWaybillProductByIdResponse struct {
	WaybillProduct domain.WaybillProduct `json:"waybill_product"`
}

func makeGetWaybillProductByIdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getWaybillProductByIdRequest)
		resp, err := s.GetWaybillProductById(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getListOfTransfersRequest struct {
	Authorization string
	Id            int64 `json:"id"`
	Limit         int   `json:"limit"`
	Offset        int   `json:"offset"`
}

type getListOfTransfersResponse struct {
	Transfers []domain.Transfer `json:"transfers"`
}

func makeGetListOfTransfersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getListOfTransfersRequest)
		resp, err := s.GetListOfTransfers(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getWaybillProductByBarcodeRequest struct {
	Authorization string
	Barcode       string `json:"barcode"`
	WaybillId     int64  `json:"waybill_id"`
}

type getWaybillProductByBarcodeResponse struct {
	Product domain.WaybillProduct `json:"product"`
	Found   bool                  `json:"found"`
}

func makeGetWaybillProductByBarcodeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getWaybillProductByBarcodeRequest)
		resp, err := s.GetWaybillProductByBarcode(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type saveOrderRequest struct {
	Authorization string
	Order         domain.Order `json:"order"`
}

type saveOrderResponse struct {
	Id int64 `json:"id"`
}

func makeSaveOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(saveOrderRequest)
		resp, err := s.SaveOrder(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getOrderRequest struct {
	Authorization string
	Id            int64 `json:"id"`
}

type getOrderResponse struct {
	Order domain.Order `json:"order"`
}

func makeGetOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderRequest)
		resp, err := s.GetOrder(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type getOrdersListRequest struct {
	Authorization string
	MerchantId    string `json:"merchant_id"`
	StockId       string `json:"stock_id"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
}

type getOrdersListResponse struct {
	Orders []domain.Order `json:"orders"`
}

func makeGetOrdersListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersListRequest)
		resp, err := s.GetOrdersList(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type syncProductsRequest struct {
	Authorization string
	MerchantId    string    `json:"merchant_id"`
	StockId       string    `json:"stock_id"`
	LastUpdate    time.Time `json:"last_update"`
}

type syncProductsResponse struct {
	Products []domain.Product `json:"products"`
}

func makeSyncProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(syncProductsRequest)
		resp, err := s.SyncProducts(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type updateWaybillRequest struct {
	Authorization string
	Waybill       domain.Waybill `json:"waybill"`
}

type updateWaybilllResponse struct {
	Message string `json:"message"`
}

func makeUpdateWaybillEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateWaybillRequest)
		resp, err := s.UpdateWaybill(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
