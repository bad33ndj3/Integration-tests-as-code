package repository

import "integrationtest/svc/cart/proto"

type CartStore struct {
	store map[string]*proto.Cart
}

func New(store map[string]*proto.Cart) *CartStore {
	return &CartStore{store: store}
}

func (c *CartStore) GetCart(userId string) (*proto.Cart, error) {
	return c.store[userId], nil
}

func (c *CartStore) AddCartItem(userId string, productId string) error {

	var exists bool
	// if item already exists, increase quantity
	for _, item := range c.store[userId].Items {
		if item.ProductId == productId {
			item.Quantity++
			exists = true
			return nil
		}
	}
	if !exists {
		c.store[userId].Items = append(c.store[userId].Items, &proto.CartItem{ProductId: productId, Quantity: 1})
	}

	return nil
}

func (c *CartStore) RemoveCartItem(userId string, productId string) error {
	for i, item := range c.store[userId].Items {
		if item.ProductId == productId {
			c.store[userId].Items = append(c.store[userId].Items[:i], c.store[userId].Items[i+1:]...)
			return nil
		}
	}
	return nil
}
