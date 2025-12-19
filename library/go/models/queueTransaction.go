package models

type QueueTransaction struct {
	ID              []byte                    `json:"id" bson:"id"`
	QueueCustomerId []byte                    `json:"queueCustomerId" bson:"queueCustomerId"`
	Details         []QueueTransactionDetails `json:"details" bson:"details"`
	Cursor          []byte                    `json:"cursor" bson:"cursor"`
	IsReserve       bool                      `json:"isReserve" bson:"isReserve"`
	CreatedAt       int64                     `json:"createdAt" bson:"createdAt"`
	UpdatedAt       int64                     `json:"updatedAt" bson:"updatedAt"`
}

type QueueTransactionDetails struct {
	QueueItemId []byte `json:"queueItemId" bson:"queueItemId"`
	Order       int    `json:"order" bson:"order"`
	IsCalled    bool   `json:"isCalled" bson:"isCalled"`
	IsDone      bool   `json:"isDone" bson:"isDone"`
	EmployeeId  []byte `json:"employeeId" bson:"employeeId"` // who called the queue
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
}
