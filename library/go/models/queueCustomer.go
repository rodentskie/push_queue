package models

type QueueCustomer struct {
	ID                  []byte `json:"id" bson:"id"`
	QueueString         string `json:"queueString" bson:"queueString"`
	ActiveTransactionId []byte `json:"activeTransactionId,omitempty" bson:"activeTransactionId,omitempty"` // Currently active transaction (must complete before next)
	TransactionSequence int64  `json:"transactionSequence" bson:"transactionSequence"`                     // Counter for ordering transactions
	TotalTransactions   int64  `json:"totalTransactions" bson:"totalTransactions"`                         // Total number of transactions (for metrics)
	TotalCompletedTxn   int64  `json:"totalCompletedTxn" bson:"totalCompletedTxn"`                         // Total completed transactions
	Cursor              []byte `json:"cursor" bson:"cursor"`
	CreatedAt           int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt           int64  `json:"updatedAt" bson:"updatedAt"`
}
