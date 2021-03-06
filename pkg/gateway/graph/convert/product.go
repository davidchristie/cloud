package convert

import (
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
	"github.com/davidchristie/cloud/pkg/product"
)

func Product(product *product.Product) *model.Product {
	if product == nil {
		return nil
	}
	return &model.Product{
		Description: product.Description,
		ID:          product.ID.String(),
		Name:        product.Name,
	}
}
