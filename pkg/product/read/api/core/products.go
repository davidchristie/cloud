package core

import (
	"context"
	"log"

	"github.com/davidchristie/cloud/pkg/entity"
)

func (c *core) Products(ctx context.Context) ([]*entity.Product, error) {
	products, err := c.ProductRepository.FindProducts(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}
