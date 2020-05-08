package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/product"
	"github.com/davidchristie/cloud/pkg/product/database"
)

type Core interface {
	Product(context.Context, string) (*product.Product, error)
	Products(context.Context) ([]*product.Product, error)
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
