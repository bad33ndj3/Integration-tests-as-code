package clientconv

import (
	context "context"
	"integrationtest/svc/product/proto"

	grpc2 "google.golang.org/grpc"
)

type ProductClientConv struct {
	server proto.ProductServiceServer
}

func NewProductClientConv(server proto.ProductServiceServer) *ProductClientConv {
	return &ProductClientConv{server: server}
}

func (p *ProductClientConv) GetProducts(ctx context.Context, in *proto.Empty, opts ...grpc2.CallOption) (*proto.ProductList, error) {
	return p.server.GetProducts(ctx, in)
}

func (p *ProductClientConv) GetProductById(ctx context.Context, in *proto.ProductId, opts ...grpc2.CallOption) (*proto.Product, error) {
	return p.server.GetProductById(ctx, in)
}

func (p *ProductClientConv) CreateProduct(ctx context.Context, in *proto.Product, opts ...grpc2.CallOption) (*proto.ProductId, error) {
	return p.server.CreateProduct(ctx, in)
}

func (p *ProductClientConv) UpdateProduct(ctx context.Context, in *proto.Product, opts ...grpc2.CallOption) (*proto.Empty, error) {
	return p.server.UpdateProduct(ctx, in)
}
