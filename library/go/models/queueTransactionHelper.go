package models

// Helper methods and business logic for sequential transaction processing

// CanBeActivated checks if a transaction can be activated based on customer's state
// Returns true only if:
// 1. Customer has no active transaction, OR
// 2. This transaction IS the active transaction (reactivation scenario)
func (qt *QueueTransaction) CanBeActivated(customer *QueueCustomer) bool {
	// If customer has no active transaction, any pending transaction can be activated
	if len(customer.ActiveTransactionId) == 0 {
		return qt.Status == TransactionStatusPending
	}

	// If this transaction is already the active one, it can be reactivated
	if bytesEqual(qt.ID, customer.ActiveTransactionId) {
		return qt.Status == TransactionStatusActive
	}

	// Otherwise, cannot activate (another transaction is active)
	return false
}

// IsNextInSequence checks if this transaction is the next one to be processed
// for the given customer
func (qt *QueueTransaction) IsNextInSequence(customer *QueueCustomer) bool {
	// The next transaction should have sequence = customer's current sequence + 1
	// OR if customer has no completed transactions, sequence should be 1
	if customer.TotalCompletedTxn == 0 {
		return qt.SequenceNumber == 1
	}
	return qt.SequenceNumber == customer.TransactionSequence
}

// AllItemsCompleted checks if all queue items in this transaction are done
func (qt *QueueTransaction) AllItemsCompleted() bool {
	if len(qt.Details) == 0 {
		return false
	}

	for _, detail := range qt.Details {
		if !detail.IsDone {
			return false
		}
	}
	return true
}

// HasPendingItems checks if transaction has any items not yet completed
func (qt *QueueTransaction) HasPendingItems() bool {
	for _, detail := range qt.Details {
		if !detail.IsDone {
			return true
		}
	}
	return false
}

// GetNextQueueItem returns the next queue item to be called (first non-called item)
func (qt *QueueTransaction) GetNextQueueItem() *QueueTransactionDetails {
	for i := range qt.Details {
		if !qt.Details[i].IsCalled {
			return &qt.Details[i]
		}
	}
	return nil
}

// Helper function to compare byte slices
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
