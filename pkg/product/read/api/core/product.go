package core

import (
	"context"
	"errors"
	"log"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/product/database"
	"github.com/google/uuid"
)

var ErrProductNotFound = errors.New("product not found")

func (c *core) Product(ctx context.Context, id string) (*entity.Product, error) {
	productID, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		return nil, ErrProductNotFound
	}
	product, err := c.ProductRepository.FindProduct(ctx, productID)
	switch err {
	case nil:
		return product, nil

	case database.ErrProductNotFound:
		return nil, ErrProductNotFound

	default:
		return nil, err
	}
}
