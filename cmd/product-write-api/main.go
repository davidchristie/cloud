package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/product/write/api"
)

func main() {
	log.Fatal(api.StartService())
}
