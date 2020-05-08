package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/customer/read/api"
)

func main() {
	log.Fatal(api.StartService())
}
