package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/order/read/api"
)

func main() {
	log.Fatal(api.StartService())
}
