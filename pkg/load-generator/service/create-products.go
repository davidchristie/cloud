package service

import (
	"fmt"
	"time"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (s *service) CreateFakeProduct() (*entity.Product, error) {
	fmt.Println("create fake product")
	return s.ProductWriteAPI.CreateProduct(fake.ProductName(), fake.Sentences(), uuid.New())
}

func (s *service) GenerateFakeProducts() {
	for i := 0; ; i++ {
		_, err := s.CreateFakeProduct()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
