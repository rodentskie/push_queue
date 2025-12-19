package models

import "encoding/json"

type Eventstore struct {
	ID        []byte          `json:"id" bson:"id"`
	Event     string          `json:"event" bson:"event"`
	EventId   int             `json:"eventId" bson:"eventId"`
	Body      json.RawMessage `json:"body" bson:"body"`
	CreatedAt int64           `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64           `json:"updatedAt" bson:"updatedAt"`
}
