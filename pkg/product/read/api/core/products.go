package core

import (
	"context"
	"log"

	"github.com/davidchristie/cloud/pkg/product"
)

func (c *core) Products(ctx context.Context) ([]*product.Product, error) {
	products, err := c.ProductRepository.FindProducts(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}
