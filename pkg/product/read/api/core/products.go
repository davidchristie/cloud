package core

import (
	"context"
	"log"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/google/uuid"
)

func (c *core) Products(ctx context.Context, ids []string) (map[uuid.UUID]*product.Product, error) {
	productIDs := []uuid.UUID{}
	for _, id := range ids {
		productID, err := uuid.Parse(id)
		if err == nil {
			productIDs = append(productIDs, productID)
		} else {
			log.Println(err)
		}
	}
	return c.ProductRepository.FindProducts(ctx, productIDs)
}
