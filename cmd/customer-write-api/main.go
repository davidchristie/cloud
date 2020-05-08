package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/customer/write/api"
)

func main() {
	log.Fatal(api.StartService())
}
