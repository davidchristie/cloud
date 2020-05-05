package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/product/worker"
)

func main() {
	log.Fatal(worker.StartService())
}
