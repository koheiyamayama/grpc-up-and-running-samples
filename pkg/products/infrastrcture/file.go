package infrastrcture

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	products "github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products"
)

type FileClientForProducts struct{}

func NewFileClientForProducts() products.ProductsRepository {
	return &FileClientForProducts{}
}

func (f *FileClientForProducts) GetProduct(ctx context.Context, params products.GetProductParams) (*products.GetProductRecord, error) {
	// TODO: database/products/products.jsonを作成する
	//       指定されたIDのproductsをjsonファイルから取得する処理を記述する
	file, err := os.Open("./database/products/products.json")
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to open database file: %w", ctx, params, err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to io.ReadAll from file: %w", ctx, params, err)
	}

	type JsonProduct struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		SalePrice     int64  `json:"sale_price"`
		OriginalPrice int64  `json:"original_price"`
		CategoryIDs   []int  `json:"category_ids"`
	}

	records := []*JsonProduct{}
	err = json.Unmarshal(b, &records)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to json.Unmarshal from file to records: %w", ctx, params, err)
	}

	var getProductRecord *products.GetProductRecord
	for _, r := range records {
		if params.ID == r.ID {
			getProductRecord = &products.GetProductRecord{
				ID:            r.ID,
				Name:          r.Name,
				SalePrice:     r.SalePrice,
				OriginalPrice: r.OriginalPrice,
			}
		}
	}

	if getProductRecord != nil {
		return getProductRecord, nil
	} else {
		// TODO: NotFoundエラーをrepositoryに定義してそれを使う
		return nil, errors.New("not found")
	}
}
