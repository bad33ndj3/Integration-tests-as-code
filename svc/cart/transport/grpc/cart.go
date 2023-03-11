package grpc

import (
	"context"
	"integrationtest/svc/cart/proto"
	productProto "integrationtest/svc/product/proto"
)

type Repository interface {
	GetCart(userId string) (*proto.Cart, error)
	AddCartItem(userId string, productId string) error
	RemoveCartItem(userId string, productId string) error
}

type Server struct {
	cartRepository Repository
	productClient  productProto.ProductServiceClient
}

func New(cartRepository Repository, productClient productProto.ProductServiceClient) *Server {
	return &Server{cartRepository: cartRepository, productClient: productClient}
}

func (s *Server) GetCart(ctx context.Context, id *proto.CartUserId) (*proto.Cart, error) {
	cart, err := s.cartRepository.GetCart(id.UserId)
	if err != nil {
		return nil, err
	}

	var total float32
	for _, item := range cart.Items {
		product, err := s.productClient.GetProductById(ctx, &productProto.ProductId{Id: item.ProductId})
		if err != nil {
			return nil, err
		}

		total += product.Price * float32(item.Quantity)
	}

	cart.TotalPrice = int32(total)

	return cart, nil
}

func (s *Server) AddCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.Empty, error) {
	return nil, s.cartRepository.AddCartItem(request.UserId, request.ProductId)
}

func (s *Server) RemoveCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.Empty, error) {
	return nil, s.cartRepository.RemoveCartItem(request.UserId, request.ProductId)
}
