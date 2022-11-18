package products

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	categoryv1 "github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/categories/v1"
	productv1 "github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/products/v1"
	productsRepo "github.com/koheiyamayama/grpc-up-and-running-samples/pkg/products/repository"
)

type ProductsServer struct {
	productClient productsRepo.ProductsRepository
}

func NewProductsServer(productRepo productsRepo.ProductsRepository) *ProductsServer {
	return &ProductsServer{
		productClient: productRepo,
	}
}

func (s *ProductsServer) GetProduct(
	ctx context.Context,
	req *connect.Request[productv1.GetProductRequest],
) (*connect.Response[productv1.GetProductResponse], error) {
	p, err := s.productClient.GetProduct(ctx, productsRepo.GetProductParams{ID: req.Msg.ProductId})
	if err != nil {
		if errors.Is(err, productsRepo.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		} else {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}
	res := connect.NewResponse(&productv1.GetProductResponse{
		Product: &productv1.Product{
			Id:            p.ID,
			Name:          p.Name,
			SalePrice:     p.SalePrice,
			OriginalPrice: p.OriginalPrice,
			Categories: func() []*categoryv1.Category {
				categories := make([]*categoryv1.Category, len(p.Categories))
				for i, c := range p.Categories {
					categories[i] = &categoryv1.Category{
						Id:   c.ID,
						Name: c.Name,
					}
				}

				return categories
			}(),
		},
	})

	return res, nil
}

func (s *ProductsServer) ListProducts(
	ctx context.Context,
	req *connect.Request[productv1.ListProductsRequest],
) (*connect.Response[productv1.ListProductsResponse], error) {
	return &connect.Response[productv1.ListProductsResponse]{}, nil
}

func (s *ProductsServer) RegisterProducts(
	ctx context.Context,
	req *connect.Request[productv1.RegisterProductsRequest],
) (*connect.Response[productv1.RegisterProductsResponse], error) {
	return &connect.Response[productv1.RegisterProductsResponse]{}, nil
}

func (s *ProductsServer) UnregisterProducts(
	ctx context.Context,
	req *connect.Request[productv1.UnregisterProductsRequest],
) (*connect.Response[productv1.UnregisterProductsResponse], error) {
	return &connect.Response[productv1.UnregisterProductsResponse]{}, nil
}
