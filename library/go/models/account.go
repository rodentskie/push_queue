package models

type Account struct {
	ID           []byte `json:"id" bson:"id"`
	Username     string `json:"username" bson:"username"`
	Password     string `json:"password" bson:"password"`
	EmployeeId   []byte `json:"employeeId" bson:"employeeId"`
	RoleId       []byte `json:"roleId" bson:"roleId"`
	Cursor       []byte `json:"cursor" bson:"cursor"`
	AccessToken  string `json:"accessToken" bson:"accessToken"`
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
	LoginCount   int    `json:"loginCount" bson:"loginCount"`
	Status       int    `json:"status" bson:"status"`
	CreatedBy    []byte `json:"createdBy" bson:"createdBy"`
	UpdatedBy    []byte `json:"updatedBy" bson:"updatedBy"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt" bson:"updatedAt"`
}
