package constant

// TransactionStatus represents the lifecycle states of a queue transaction
const (
	TransactionStatusActive    = 1
	TransactionStatusCompleted = 2
)

// QueueItemStatus represents the lifecycle states of individual queue items within a transaction
const (
	QueueItemStatusPending   = 0
	QueueItemStatusCalled    = 1
	QueueItemStatusCompleted = 2
	QueueItemStatusSkipped   = 3
)
