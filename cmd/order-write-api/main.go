package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/order/write/api"
)

func main() {
	log.Fatal(api.StartService())
}
