package utility

import (
	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/davidchristie/cloud/pkg/gateway/graph/model"
)

func ConvertProductToModel(product *entity.Product) *model.Product {
	return &model.Product{
		Description: product.Description,
		ID:          product.ID.String(),
		Name:        product.Name,
	}
}
