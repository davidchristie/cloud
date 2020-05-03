package service

import (
	"fmt"
	"time"

	"github.com/davidchristie/cloud/pkg/customer-write-api/client"
	"github.com/icrowley/fake"
)

func CreateCustomers() {
	customerWriteAPI := client.NewClient()

	for i := 0; ; i++ {
		fmt.Println("create customer")
		_, err := customerWriteAPI.CreateCustomer(fake.FirstName(), fake.LastName())
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
