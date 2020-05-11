package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/search/worker"
)

func main() {
	log.Fatal(worker.StartService())
}
