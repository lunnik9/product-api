package sources

import (
	"context"
	"encoding/json"
	"net/http"

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

	r := mux.NewRouter()

	r.Handle("/merch/login", login).Methods("POST")
	r.Handle("/merch/refresh", getRefreshToken).Methods("POST")

	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body loginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeGetRefreshTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body getRefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
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
