package models

type QueueItem struct {
	ID          []byte `json:"id" bson:"id"`
	TeamId      []byte `json:"teamId" bson:"teamId"` // Team handles this queue item (Team -> Department)
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status      int    `json:"status" bson:"status"`
	Cursor      []byte `json:"cursor" bson:"cursor"`
	ApprovedBy  []byte `json:"approvedBy" bson:"approvedBy"`
	CreatedBy   []byte `json:"createdBy" bson:"createdBy"`
	UpdatedBy   []byte `json:"updatedBy" bson:"updatedBy"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
}
