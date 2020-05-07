package core

import (
	"context"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/product/database"
)

type Core interface {
	Product(context.Context, string) (*entity.Product, error)
	Products(context.Context) ([]*entity.Product, error)
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
