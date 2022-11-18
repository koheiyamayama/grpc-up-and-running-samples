package products

import (
	"context"

	"github.com/koheiyamayama/grpc-up-and-running-samples/pkg/categories"
)

type ProductsRepository interface {
	GetProduct(ctx context.Context, params GetProductParams) (*GetProductRecord, error)
}

type GetProductParams struct {
	ID string
}

type GetProductRecord struct {
	ID            string
	Name          string
	OriginalPrice int64
	SalePrice     int64
	Categories    []categories.Category
}
