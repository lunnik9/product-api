package sources

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	pe "github.com/lunnik9/product-api/product_errors"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/lunnik9/product-api/product_errors"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func MakeHandler(ss Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	login := kithttp.NewServer(
		makeLoginEndpoint(ss),
		decodeLoginRequest,
		encodeResponse,
		opts...,
	)

	getRefreshToken := kithttp.NewServer(
		makeGetRefreshTokenEndpoint(ss),
		decodeGetRefreshTokenRequest,
		encodeResponse,
		opts...,
	)

	listMerchantStocks := kithttp.NewServer(
		makeListMerchantStocksEndpoint(ss),
		decodeListMerchantStocksRequest,
		encodeResponse,
		opts...,
	)

	getProductById := kithttp.NewServer(
		makeGetProductByIdEndpoint(ss),
		decodeGetProductByIdRequset,
		encodeResponse,
		opts...,
	)

	createProduct := kithttp.NewServer(
		makeCreateProductEndpoint(ss),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	)

	updateProduct := kithttp.NewServer(
		makeUpdateProductEndpoint(ss),
		decodeUpdateProductRequest,
		encodeResponse,
		opts...,
	)

	deleteProduct := kithttp.NewServer(
		makeDeleteProductEndpoint(ss),
		decodeDeleteProductRequest,
		encodeResponse,
		opts...,
	)

	filterProducts := kithttp.NewServer(
		makeFilterProductsEndpoint(ss),
		decodeFilterProductsRequest,
		encodeResponse,
		opts...,
	)

	getListOfCashBoxes := kithttp.NewServer(
		makeGetListOfCashBoxesEndpoint(ss),
		decodeHetListOfCashboxesRequest,
		encodeResponse,
		opts...,
	)

	getCategoryById := kithttp.NewServer(
		makeGetCategoryByIdEndpoint(ss),
		decodeGetCategoryByIdRequset,
		encodeResponse,
		opts...,
	)

	createCategory := kithttp.NewServer(
		makeCreateCategoryEndpoint(ss),
		decodeCreateCategoryRequest,
		encodeResponse,
		opts...,
	)

	updateCategory := kithttp.NewServer(
		makeUpdateCategoryEndpoint(ss),
		decodeUpdateCategoryRequest,
		encodeResponse,
		opts...,
	)

	deleteCategory := kithttp.NewServer(
		makeDeleteCategoryEndpoint(ss),
		decodeDeleteCategoryRequest,
		encodeResponse,
		opts...,
	)

	filterCategories := kithttp.NewServer(
		makeFilterCategoriesEndpoint(ss),
		decodeFilterCategoriesRequest,
		encodeResponse,
		opts...,
	)

	mDeleteProducts := kithttp.NewServer(
		makeMDeleteProductsEndpoint(ss),
		decodeMDeleteProductsRequest,
		encodeResponse,
		opts...,
	)

	createWaybill := kithttp.NewServer(
		makeCreateWaybillEndpoint(ss),
		decodeCreateWaybillRequest,
		encodeResponse,
		opts...,
	)

	conductWaybill := kithttp.NewServer(
		makeConductWaybillEndpoint(ss),
		decodeConductWaybillRequest,
		encodeResponse,
		opts...,
	)

	rollbackWaybill := kithttp.NewServer(
		makeRollbackWaybillEndpoint(ss),
		decodeRollbackWaybillRequest,
		encodeResponse,
		opts...,
	)

	deleteWaybill := kithttp.NewServer(
		makeDeleteWaybillEndpoint(ss),
		decodeDeleteWaybillRequest,
		encodeResponse,
		opts...,
	)

	filterWaybills := kithttp.NewServer(
		makeFilterWaybillsEndpoint(ss),
		decodeFilterWaybillsRequest,
		encodeResponse,
		opts...,
	)

	createWaybillProduct := kithttp.NewServer(
		makeCreateWaybillProductEndpoint(ss),
		decodeCreateWaybillProductRequest,
		encodeResponse,
		opts...,
	)

	updateWaybillProduct := kithttp.NewServer(
		makeUpdateWaybillProductEndpoint(ss),
		decodeUpdateWaybillProductRequest,
		encodeResponse,
		opts...,
	)

	deleteWaybillProduct := kithttp.NewServer(
		makeDeleteWaybillProductEndpoint(ss),
		decodeDeleteWaybillProductRequest,
		encodeResponse,
		opts...,
	)

	getWaybillProductsList := kithttp.NewServer(
		makeGetListOfWaybillProductsEndpoint(ss),
		decodeGetWaybillProductsList,
		encodeResponse,
		opts...,
	)

	getWaybillById := kithttp.NewServer(
		makeGetWaybillByIdEndpoint(ss),
		decodeGetWaybillByIdRequest,
		encodeResponse,
		opts...,
	)

	getWaybillProductById := kithttp.NewServer(
		makeGetWaybillProductByIdEndpoint(ss),
		decodeGetWaybillProductByIdRequest,
		encodeResponse,
		opts...,
	)

	getWaybillProductByBarcode := kithttp.NewServer(
		makeGetWaybillProductByBarcodeEndpoint(ss),
		decodeGetWaybillProductByProductIdRequest,
		encodeResponse,
		opts...,
	)

	getListOfTransfers := kithttp.NewServer(
		makeGetListOfTransfersEndpoint(ss),
		decodeGetListOfTransfersRequest,
		encodeResponse,
		opts...,
	)

	saveOrder := kithttp.NewServer(
		makeSaveOrderEndpoint(ss),
		decodeSaveOrderRequest,
		encodeResponse,
		opts...,
	)

	getOrder := kithttp.NewServer(
		makeGetOrderEndpoint(ss),
		decodeGetOrderRequest,
		encodeResponse,
		opts...,
	)

	getOrdersList := kithttp.NewServer(
		makeGetOrdersListEndpoint(ss),
		decodeGetOrdersListRequest,
		encodeResponse,
		opts...,
	)

	syncProducts := kithttp.NewServer(
		makeSyncProductsEndpoint(ss),
		decodeSyncProductsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/merch/login", login).Methods("POST")
	r.Handle("/merch/refresh", getRefreshToken).Methods("POST")

	r.Handle("/stocks/list/{merchant_id}", listMerchantStocks).Methods("GET")

	r.Handle("/cashbox/list", getListOfCashBoxes).Methods("POST")

	r.Handle("/product/{product_id}", getProductById).Methods("GET")
	r.Handle("/product", createProduct).Methods("POST")
	r.Handle("/product", updateProduct).Methods("PUT")
	r.Handle("/product/{product_id}", deleteProduct).Methods("DELETE")
	r.Handle("/product/filter", filterProducts).Methods("POST")
	r.Handle("/product/mdelete", mDeleteProducts).Methods("POST")
	r.Handle("/product/sync", syncProducts).Methods("POST")

	r.Handle("/transfer/get_list", getListOfTransfers).Methods("POST")

	r.Handle("/category/{category_id}", getCategoryById).Methods("GET")
	r.Handle("/category", createCategory).Methods("POST")
	r.Handle("/category", updateCategory).Methods("PUT")
	r.Handle("/category/{category_id}", deleteCategory).Methods("DELETE")
	r.Handle("/category/filter", filterCategories).Methods("POST")

	r.Handle("/waybill/create", createWaybill).Methods("POST")
	r.Handle("/waybill/conduct/{waybill_id}", conductWaybill).Methods("GET")
	r.Handle("/waybill/rollback/{waybill_id}", rollbackWaybill).Methods("GET")
	r.Handle("/waybill/{waybill_id}", deleteWaybill).Methods("DELETE")
	r.Handle("/waybill/{waybill_id}", getWaybillById).Methods("GET")
	r.Handle("/waybill/filter", filterWaybills).Methods("POST")

	r.Handle("/waybill/product", createWaybillProduct).Methods("POST")
	r.Handle("/waybill/product", updateWaybillProduct).Methods("PUT")
	r.Handle("/waybill/product/{product_id}", deleteWaybillProduct).Methods("DELETE")
	r.Handle("/waybill/product/{product_id}", getWaybillProductById).Methods("GET")
	r.Handle("/waybill/product/get_list", getWaybillProductsList).Methods("POST")
	r.Handle("/waybill/product/get", getWaybillProductByBarcode).Methods("POST")

	r.Handle("/order/{order_id}", getOrder).Methods("GET")
	r.Handle("/order/list", getOrdersList).Methods("POST")
	r.Handle("/order/create", saveOrder).Methods("POST")

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body loginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}
	return body, nil
}

func decodeGetRefreshTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getRefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}
	return body, nil
}

func decodeListMerchantStocksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	merchantId, ok := vars["merchant_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return listMerchantStocksRequest{token, merchantId}, nil
}

func decodeGetProductByIdRequset(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["product_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return getProductByIdRequest{token, productId}, nil
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createProductRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body updateProductRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["product_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return deleteProductRequest{token, productId}, nil
}

func decodeFilterProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body filterProductsRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeHetListOfCashboxesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getListOfCashBoxesRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeGetCategoryByIdRequset(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["category_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	categoryId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return getCategoryByIdRequest{token, categoryId}, nil
}

func decodeCreateCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createCategoryRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeUpdateCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body updateCategoryRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeDeleteCategoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["category_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	categoryId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return deleteCategoryRequest{token, categoryId}, nil
}

func decodeFilterCategoriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body filterCategoriesRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeMDeleteProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body mDeleteProductsRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeCreateWaybillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createWaybillRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeConductWaybillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["waybill_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return conductWaybillRequest{token, productId}, nil
}

func decodeRollbackWaybillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["waybill_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return rollbackWaybillRequest{token, productId}, nil
}

func decodeDeleteWaybillRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["waybill_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return deleteWaybillRequest{token, productId}, nil
}

func decodeFilterWaybillsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body filterWaybillsRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeCreateWaybillProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createWaybillProductRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeUpdateWaybillProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body updateWaybillProductRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeDeleteWaybillProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["product_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return deleteWaybillProductRequest{token, productId}, nil
}

func decodeGetWaybillProductsList(_ context.Context, r *http.Request) (interface{}, error) {
	var body getListOfWaybillProductsRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeGetWaybillByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["waybill_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return getWaybillByIdRequest{token, productId}, nil
}

func decodeGetWaybillProductByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	productIdString, ok := vars["product_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	productId, err := strconv.ParseInt(productIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return getWaybillProductByIdRequest{token, productId}, nil
}

func decodeGetOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	orderIdString, ok := vars["order_id"]
	if !ok {
		return nil, pe.New(409, "no merch id provided")
	}

	orderId, err := strconv.ParseInt(orderIdString, 10, 64)
	if err != nil {
		return nil, pe.New(409, "no product id provided")
	}

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	return getOrderRequest{token, orderId}, nil
}

func decodeGetListOfTransfersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getListOfTransfersRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeGetWaybillProductByProductIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getWaybillProductByBarcodeRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeSaveOrderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body saveOrderRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func decodeGetOrdersListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getOrdersListRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}
func decodeSyncProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body syncProductsRequest

	token, err := getAuthorizationToken(r)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, pe.New(409, err.Error())
	}

	body.Authorization = token
	return body, nil
}

func getAuthorizationToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")

	authorizationSlice := strings.Split(authorization, " ")

	if strings.ToLower(authorizationSlice[0]) != "bearer" {
		return "", pe.New(401, "no authorization token provided")
	}

	return authorizationSlice[1], nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode product_errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	respErr := err.(product_errors.ProductError)
	w.WriteHeader(respErr.StatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encodeErr := json.NewEncoder(w).Encode(respErr)
	if encodeErr != nil {
		Logger.Error(err)
	}
}
