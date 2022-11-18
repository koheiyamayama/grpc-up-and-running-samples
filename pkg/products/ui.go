package products

import (
	"context"

	"github.com/bufbuild/connect-go"
	productv1 "github.com/koheiyamayama/grpc-up-and-running-samples/gen/proto/products/v1"
)

type ProductsServer struct {
	ProductDAO ProductsRepository
}

func NewProductsServer(dao ProductsRepository) *ProductsServer {
	return &ProductsServer{
		ProductDAO: dao,
	}
}

func (s *ProductsServer) GetProduct(
	ctx context.Context,
	req *connect.Request[productv1.GetProductRequest],
) (*connect.Response[productv1.GetProductResponse], error) {
	p, err := s.ProductDAO.GetProduct(ctx, GetProductParams{ID: req.Msg.ProductId})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&productv1.GetProductResponse{
		Product: &productv1.Product{
			Id:            p.ID,
			Name:          p.Name,
			SalePrice:     p.SalePrice,
			OriginalPrice: p.OriginalPrice,
			Categories: []*productv1.Category{
				{
					Id:   "categoryOne",
					Name: "アニメ",
				},
				{
					Id:   "categoryTwo",
					Name: "少年漫画",
				},
			},
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
