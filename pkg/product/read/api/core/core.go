package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/davidchristie/cloud/pkg/product/database"
	"github.com/google/uuid"
)

type Core interface {
	Product(context.Context, string) (*product.Product, error)
	Products(context.Context, []string) (map[uuid.UUID]*product.Product, error)
}

type core struct {
	ProductRepository database.ProductRepository
}

func NewCore() Core {
	db := database.Connect()
	return &core{
		ProductRepository: database.NewProductRepository(db),
	}
}
