package service

import (
	"fmt"
	"time"

	"github.com/davidchristie/cloud/pkg/product-write-api/client"
	"github.com/icrowley/fake"
)

func CreateProducts() {
	productWriteAPI := client.NewClient()

	for i := 0; ; i++ {
		fmt.Println("create product")
		_, err := productWriteAPI.CreateProduct(fake.ProductName(), fake.Sentences())
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
