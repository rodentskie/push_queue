package models

// TransactionStatus represents the lifecycle states of a queue transaction
const (
	TransactionStatusPending   = 0 // Created but not yet active
	TransactionStatusActive    = 1 // Currently being processed (customer called)
	TransactionStatusCompleted = 2 // All queue items in transaction are done
	TransactionStatusCancelled = 3 // Transaction was cancelled
	TransactionStatusSkipped   = 4 // Customer skipped/no-show
)

type QueueTransaction struct {
	ID              []byte                    `json:"id" bson:"id"`
	QueueCustomerId []byte                    `json:"queueCustomerId" bson:"queueCustomerId"`
	SequenceNumber  int64                     `json:"sequenceNumber" bson:"sequenceNumber"` // Global sequence for this customer
	DepartmentId    []byte                    `json:"departmentId" bson:"departmentId"`     // Primary department for this transaction
	Status          int                       `json:"status" bson:"status"`                 // TransactionStatus constants
	Details         []QueueTransactionDetails `json:"details" bson:"details"`
	Cursor          []byte                    `json:"cursor" bson:"cursor"`
	IsReserve       bool                      `json:"isReserve" bson:"isReserve"`
	ActivatedAt     int64                     `json:"activatedAt,omitempty" bson:"activatedAt,omitempty"` // When transaction became active
	CompletedAt     int64                     `json:"completedAt,omitempty" bson:"completedAt,omitempty"` // When all items were completed
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
