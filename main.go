package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lunnik9/product-api/config"
	"github.com/lunnik9/product-api/sources"
	"github.com/lunnik9/product-api/sources/db"
	"github.com/lunnik9/product-api/sources/merch_repo"
	"github.com/lunnik9/product-api/sources/product_repo"
	"github.com/lunnik9/product-api/sources/waybill_repo"
)

func main() {
	err := config.GetConfigs()
	if err != nil {
		panic(err)
	}

	port := getPort()
	fmt.Println("Listening on", port)

	httpAddr := flag.String("http.addr", ":"+port, "HTTP listen address only port :"+port)

	pgConn, err := db.Connect(config.AllConfigs.Postgres.Url)
	if err != nil {
		panic(err)
	}

	var (
		//logger     log.Logger
		mr = merch_repo.NewMerchPostgres(pgConn)
		pr = product_repo.NewProductPostgres(pgConn)
		wr = waybill_repo.NewWaybillPostgres(pgConn)
		//httpLogger = log.With(logger, "component", "http")
	)

	service := sources.NewService(&mr, &pr, &wr)

	mux := http.NewServeMux()
	mux.Handle("/", sources.MakeHandler(service, log.NewLogfmtLogger(os.Stderr)))

	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()

	fmt.Println("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = config.AllConfigs.Http.Port
	}

	return port
}
