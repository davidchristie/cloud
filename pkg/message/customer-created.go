package message

import "github.com/davidchristie/cloud/pkg/entity"

type CustomerCreatedEvent struct {
	Data     *entity.Customer `json:"data"`
	Metadata *EventMetadata   `json:"metadata"`
}
