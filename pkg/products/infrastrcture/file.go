package infrastrcture

import (
	"context"

	products "github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products"
)

type FileClientForProducts struct{}

func NewFileClientForProducts() products.ProductsRepository {
	return &FileClientForProducts{}
}

func (f *FileClientForProducts) GetProduct(ctx context.Context, params products.GetProductParams) (*products.GetProductRecord, error) {
	// TODO: database/products/products.jsonを作成する
	//       指定されたIDのproductsをjsonファイルから取得する処理を記述する
	return &products.GetProductRecord{}, nil
}
