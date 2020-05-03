package service

func Start() {
	go CreateCustomers()
	CreateProducts()
}
