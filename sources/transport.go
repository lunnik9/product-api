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

	r := mux.NewRouter()

	r.Handle("/merch/login", login).Methods("POST")
	r.Handle("/merch/refresh", getRefreshToken).Methods("POST")

	r.Handle("/stocks/list/{merchant_id}", listMerchantStocks).Methods("GET")

	r.Handle("/product/{product_id}", getProductById).Methods("GET")
	r.Handle("/product/", createProduct).Methods("POST")
	r.Handle("/product/", updateProduct).Methods("PUT")
	r.Handle("/product/{product_id}", deleteProduct).Methods("DELETE")
	r.Handle("/product/filter", filterProducts).Methods("POST")

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
