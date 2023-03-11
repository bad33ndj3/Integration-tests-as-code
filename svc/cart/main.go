package main

import (
	"integrationtest/svc/cart/internal/repository"
	"integrationtest/svc/cart/proto"
	"integrationtest/svc/cart/transport/grpc"
	productProto "integrationtest/svc/product/proto"
	"log"

	grpc2 "google.golang.org/grpc"
)

func main() {
	cartRepository := repository.New(make(map[string]*proto.Cart))

	conn := &grpc2.ClientConn{} // should normally be a real connection
	productServiceClient := productProto.NewProductServiceClient(conn)

	cartServer := grpc.New(cartRepository, productServiceClient)

	log.Fatal(cartServer)
	//cartServer.ListenAnqServe() // should normally be called here.
}
