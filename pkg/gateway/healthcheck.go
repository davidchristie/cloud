package gateway

import (
	"log"
	"time"
)

func WaitUntilHealthy() {
	log.Println("waiting for gateway to become healthy")
	gateway := NewClient()
	query := ""
	for {
		_, err := gateway.Customers(&query)
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(5 * time.Second)
	}
}
