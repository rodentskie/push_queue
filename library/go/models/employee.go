package models

type Employee struct {
	ID            []byte `json:"id" bson:"id"`
	FirstName     string `json:"firstName" bson:"firstName"`
	LastName      string `json:"lastName" bson:"lastName"`
	BirthDate     int64  `json:"birthDate" bson:"birthDate"`
	ContactNumber string `json:"contactNumber" bson:"contactNumber"`
	TeamId        []byte `json:"teamId" bson:"teamId"`
	Status        int    `json:"status" bson:"status"`
	Cursor        []byte `json:"cursor" bson:"cursor"`
	ImageUrl      string `json:"imageUrl" bson:"imageUrl"`
	CreatedBy     []byte `json:"createdBy" bson:"createdBy"`
	UpdatedBy     []byte `json:"updatedBy" bson:"updatedBy"`
	CreatedAt     int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt" bson:"updatedAt"`
}
