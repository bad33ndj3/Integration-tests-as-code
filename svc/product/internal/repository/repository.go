package repository

import (
	"integrationtest/svc/product/proto"

	"github.com/google/uuid"
)

type ProductStore struct {
	store map[string]*proto.Product
}

func New(store map[string]*proto.Product) *ProductStore {
	return &ProductStore{store: store}
}

func (p *ProductStore) GetProducts() ([]*proto.Product, error) {
	var products []*proto.Product
	for s := range p.store {
		products = append(products, p.store[s])
	}
	return products, nil
}

func (p *ProductStore) GetProductById(id string) (*proto.Product, error) {
	return p.store[id], nil
}

func (p *ProductStore) CreateProduct(product *proto.Product) (string, error) {
	if product.Id == "" {
		product.Id = uuid.Must(uuid.NewRandom()).String()
	}
	p.store[product.Id] = product

	return product.Id, nil
}

func (p *ProductStore) UpdateProduct(product *proto.Product) error {
	p.store[product.Id] = product
	return nil
}
