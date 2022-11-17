package products

import "context"

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
	Categories    []CategoryRecord
}

type CategoryRecord struct {
	ID   string
	Name string
}
