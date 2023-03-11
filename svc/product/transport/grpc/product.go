package grpc

import (
	context "context"
	"integrationtest/svc/product/proto"
)

// Repository is the interface for the product repository.
type Repository interface {
	GetProducts() ([]*proto.Product, error)
	GetProductById(id string) (*proto.Product, error)
	CreateProduct(product *proto.Product) (string, error)
	UpdateProduct(product *proto.Product) error
}

type Server struct {
	productRepository Repository
}

func New(productRepository Repository) *Server {
	return &Server{productRepository: productRepository}
}

func (s *Server) GetProducts(ctx context.Context, empty *proto.Empty) (*proto.ProductList, error) {
	products, err := s.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	return &proto.ProductList{Products: products}, nil
}

func (s *Server) GetProductById(ctx context.Context, id *proto.ProductId) (*proto.Product, error) {
	product, err := s.productRepository.GetProductById(id.Id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Server) CreateProduct(ctx context.Context, product *proto.Product) (*proto.ProductId, error) {
	id, err := s.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &proto.ProductId{Id: id}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, product *proto.Product) (*proto.Empty, error) {
	err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
