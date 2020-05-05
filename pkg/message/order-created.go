package message

import "github.com/davidchristie/cloud/pkg/order"

type OrderCreatedEvent struct {
	Data     *order.Order   `json:"data"`
	Metadata *EventMetadata `json:"metadata"`
}
