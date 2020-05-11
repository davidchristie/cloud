package main

import (
	"log"

	"github.com/davidchristie/cloud/pkg/search/api"
)

func main() {
	log.Fatal(api.StartService())
}
