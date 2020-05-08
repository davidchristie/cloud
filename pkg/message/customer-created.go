package message

import "github.com/davidchristie/cloud/pkg/customer"

type CustomerCreatedEvent struct {
	Data     *customer.Customer `json:"data"`
	Metadata *EventMetadata     `json:"metadata"`
}
