package message

import "github.com/davidchristie/cloud/pkg/product"

type ProductCreatedEvent struct {
	Data     *product.Product `json:"data"`
	Metadata *EventMetadata   `json:"metadata"`
}
