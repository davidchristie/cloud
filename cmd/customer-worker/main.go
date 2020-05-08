package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/customer/worker"
)

func main() {
	log.Fatal(worker.StartService())
}
