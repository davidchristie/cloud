package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/load-generator/service"
)

func main() {
	log.Fatal(service.StartService())
}
