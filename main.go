package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/lunnik9/product-api/src"
	"github.com/lunnik9/product-api/src/db"
	"github.com/lunnik9/product-api/src/merch_repo"
)

var url = "postgres://pnumlsyvxztrfm:ee24c557c61258df433cfc825ea7e389ef53c907cb43195366c78f73d3c2acf4@ec2-34-252-251-16.eu-west-1.compute.amazonaws.com:5432/d1dlpo67q6hl95" //todo: change to normal config

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Listening on", port)
	httpAddr := flag.String("http.addr", ":"+port, "HTTP listen address only port :"+port)

	pgConn, err := db.Connect(url)
	if err != nil {
		panic(err)
	}

	var (
		//logger     log.Logger
		mr = merch_repo.NewMerchPostgres(pgConn)
		//httpLogger = log.With(logger, "component", "http")
	)

	service := src.NewService(&mr)

	mux := http.NewServeMux()
	mux.Handle("/", src.MakeHandler(service))

	http.Handle("/", mux)

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()

	fmt.Println("terminated", <-errs)

}
