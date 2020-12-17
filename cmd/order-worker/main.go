package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/order/worker"
)

func main() {
	log.Fatal(worker.StartService())
}
