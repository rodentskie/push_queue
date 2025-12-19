package models

type Projection struct {
	ID        []byte `json:"id" bson:"id"`
	LastEvent []byte `json:"lastEvent" bson:"lastEvent"` // eventstore ID []byte field
	EventId   int    `json:"eventId" bson:"eventId"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}
