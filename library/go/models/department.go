package models

type Department struct {
	ID           []byte   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Description  string   `json:"description" bson:"description"`
	Abbreviation string   `json:"abbreviation" bson:"abbreviation"`
	EmployeeId   [][]byte `json:"employeeId,omitempty" bson:"employeeId,omitempty"` // array of byte; a department might have multiple dept head
	Status       int      `json:"status" bson:"status"`
	Cursor       []byte   `json:"cursor" bson:"cursor"`
	CreatedBy    []byte   `json:"createdBy" bson:"createdBy"`
	UpdatedBy    []byte   `json:"updatedBy" bson:"updatedBy"`
	CreatedAt    int64    `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64    `json:"updatedAt" bson:"updatedAt"`
}
