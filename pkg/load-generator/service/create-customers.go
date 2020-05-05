package service

import (
	"fmt"
	"time"

	"github.com/davidchristie/cloud/pkg/entity"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (s *service) CreateFakeCustomer() (*entity.Customer, error) {
	fmt.Println("create fake customer")
	return s.CustomerWriteAPI.CreateCustomer(fake.FirstName(), fake.LastName(), uuid.New())
}

func (s *service) GenerateFakeCustomers() {
	for i := 0; ; i++ {
		_, err := s.CreateFakeCustomer()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
