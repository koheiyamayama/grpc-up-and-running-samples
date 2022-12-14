package main

import (
	"fmt"
	"net/http"

	"github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/products/v1/productv1connect"
	"github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products"
	productsInfra "github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products/infrastrcture"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	dao := productsInfra.NewFileClientForProducts()
	productsServer := products.NewProductsServer(dao)
	mux := http.NewServeMux()
	path, handler := productv1connect.NewProductServiceHandler(productsServer)
	mux.Handle(path, handler)
	fmt.Println("listen 8080")
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
