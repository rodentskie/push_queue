package models

type Action struct {
	ID          []byte `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	ModuleId    []byte `json:"moduleId" bson:"moduleId"`
	Cursor      []byte `json:"cursor" bson:"cursor"`
	Status      int    `json:"status" bson:"status"`
	CreatedBy   []byte `json:"createdBy" bson:"createdBy"`
	UpdatedBy   []byte `json:"updatedBy" bson:"updatedBy"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
}
