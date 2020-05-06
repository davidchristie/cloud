package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/gateway"
)

func main() {
	log.Fatal(gateway.StartService())
}
