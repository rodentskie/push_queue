package models

type QueueTransaction struct {
	ID              []byte                    `json:"id" bson:"id"`
	QueueCustomerId []byte                    `json:"queueCustomerId" bson:"queueCustomerId"`
	DepartmentId    []byte                    `json:"departmentId" bson:"departmentId"`
	Details         []QueueTransactionDetails `json:"details" bson:"details"`
	Cursor          []byte                    `json:"cursor" bson:"cursor"`
	Status          int                       `json:"status" bson:"status"`
	IsReserve       bool                      `json:"isReserve" bson:"isReserve"`
	CreatedAt       int64                     `json:"createdAt" bson:"createdAt"`
	UpdatedAt       int64                     `json:"updatedAt" bson:"updatedAt"`
}

type QueueTransactionDetails struct {
	QueueItemId []byte `json:"queueItemId" bson:"queueItemId"`
	Order       int    `json:"order" bson:"order"`
	Status      int    `json:"status" bson:"status"`
	EmployeeId  []byte `json:"employeeId" bson:"employeeId"` // who called the queue
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
}
