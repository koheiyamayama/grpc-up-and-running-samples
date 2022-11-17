package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	productv1 "github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/products/v1"
	"github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/products/v1/productv1connect"
)

func main() {
	client := productv1connect.NewProductServiceClient(
		http.DefaultClient,
		"http://localhost:50001",
	)

	res, err := client.GetProduct(
		context.Background(),
		connect.NewRequest(&productv1.GetProductRequest{
			ProductId: "gussan",
		}),
	)

	if err != nil {
		if connectErr := new(connect.Error); errors.As(err, &connectErr) {
			fmt.Println(connectErr.Message())
			fmt.Println(connectErr.Details())
			fmt.Println(connectErr.Code())
		}
	}

	log.Println(res)
}
