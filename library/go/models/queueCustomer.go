package models

type QueueCustomer struct {
	ID          []byte `json:"id" bson:"id"`
	QueueString string `json:"queueString" bson:"queueString"`
	Cursor      []byte `json:"cursor" bson:"cursor"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
}
