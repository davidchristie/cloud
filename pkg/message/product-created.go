package message

import "github.com/davidchristie/cloud/pkg/entity"

type ProductCreatedEvent struct {
	Data     *entity.Product `json:"data"`
	Metadata *EventMetadata  `json:"metadata"`
}
