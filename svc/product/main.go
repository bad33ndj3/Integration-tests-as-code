package main

import (
	"integrationtest/svc/product/internal/repository"
	"integrationtest/svc/product/proto"
	"integrationtest/svc/product/transport/grpc"
	"log"
)

func main() {
	productRepository := repository.New(make(map[string]*proto.Product))
	server := grpc.New(productRepository)

	log.Fatal(server)
	//server.ListenAnqServe() // should normally be called here.
}
