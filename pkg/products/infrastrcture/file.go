package infrastrcture

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/koheiyamayama/grpc-up-and-running-samples/pkg/categories"
	products "github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products"
)

type FileClientForProducts struct{}

func NewFileClientForProducts() products.ProductsRepository {
	return &FileClientForProducts{}
}

type jsonProduct struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	SalePrice     int64    `json:"sale_price"`
	OriginalPrice int64    `json:"original_price"`
	CategoryIDs   []string `json:"category_ids"`
}

type jsonCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (f *FileClientForProducts) GetProduct(ctx context.Context, params products.GetProductParams) (*products.GetProductRecord, error) {
	file, err := os.Open("./database/products/products.json")
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to open database file: %w", ctx, params, err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to io.ReadAll from file: %w", ctx, params, err)
	}

	records := []*jsonProduct{}
	err = json.Unmarshal(b, &records)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to json.Unmarshal from file to records: %w", ctx, params, err)
	}

	var getProductRecord *products.GetProductRecord
	for _, r := range records {
		if params.ID == r.ID {
			categoryRecords, err := f.listCategoriesByIDs(r.CategoryIDs)
			if err != nil {
				return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#GetProduct(ctx: %v, params: %v): failed to f.listCategoriesByIds(ids: %v): %w", ctx, params, r.CategoryIDs, err)
			}

			getProductRecord = &products.GetProductRecord{
				ID:            r.ID,
				Name:          r.Name,
				SalePrice:     r.SalePrice,
				OriginalPrice: r.OriginalPrice,
				Categories:    categoryRecords,
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

func (f *FileClientForProducts) listCategoriesByIDs(ids []string) ([]categories.Category, error) {
	file, err := os.Open("./database/categories/categories.json")
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#listCategoriesByIDs(ids: %v): failed to open database file: %w", ids, err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#listCategoriesByIDs(ids: %v): failed to io.ReadAll from file: %w", ids, err)
	}

	records := []jsonCategory{}
	err = json.Unmarshal(b, &records)
	if err != nil {
		return nil, fmt.Errorf("products/infrastructure/file.go: FileClientForProducts#listCategoriesByIDs(ids: %v): failed to json.Unmarshal from file to json: %w", ids, err)
	}

	var categoryRecords []categories.Category
	for _, r := range records {
		for _, id := range ids {
			if id == r.ID {
				categoryRecords = append(categoryRecords, categories.Category{
					ID:   r.ID,
					Name: r.Name,
				})
			}
		}
	}

	return categoryRecords, nil
}
