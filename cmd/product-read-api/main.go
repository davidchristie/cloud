package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/product/read/api"
)

func main() {
	log.Fatal(api.StartService())
}
