package tests

import (
	"context"
	cartProto "integrationtest/svc/cart/proto"
	cartGrpc "integrationtest/svc/cart/transport/grpc"
	cartRepoMock "integrationtest/svc/cart/transport/grpc/mock"
	productProto "integrationtest/svc/product/proto"
	productGrpc "integrationtest/svc/product/transport/grpc"
	productRepoMock "integrationtest/svc/product/transport/grpc/mock"
	"integrationtest/tests/internal/clientconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AddToCartTestSuite struct {
	suite.Suite
}

func (s *AddToCartTestSuite) TestGetCart() {
	mockProductStore := &productRepoMock.Repository{}
	mockProductStore.On("GetProductById", "1").Return(&productProto.Product{
		Id:          "1",
		Name:        "test",
		Description: "test",
		Price:       5,
	}, nil)
	productService := productGrpc.New(mockProductStore)

	mockCartStore := &cartRepoMock.Repository{}
	mockCartStore.On("GetCart", "1").Return(&cartProto.Cart{
		UserId: "1",
		Items: []*cartProto.CartItem{
			{
				ProductId: "1",
				Quantity:  2,
			},
		},
	}, nil)

	cartService := cartGrpc.New(mockCartStore, clientconv.NewProductClientConv(productService))

	ctx := context.Background()
	cart, err := cartService.GetCart(ctx, &cartProto.CartUserId{UserId: "1"})
	s.Require().NoError(err)
	s.Equal(int(10), int(cart.TotalPrice))
}

func TestAddToCartTestSuite(t *testing.T) {
	suite.Run(t, new(AddToCartTestSuite))
}
