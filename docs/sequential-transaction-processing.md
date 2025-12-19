# Sequential Transaction Processing - Design Documentation

## Problem Statement
Customers may have multiple transactions across different departments (e.g., Finance for tuition payment, then Registrar for transcript). The system must ensure that:
- **Only ONE transaction can be active per customer at any time**
- **Transactions must be processed in sequential order** (first-come, first-served)
- **Next transaction cannot be called until current transaction is fully completed**

## Solution Design

### 1. Schema Changes

#### QueueCustomer
Added fields to track transaction state:
- `ActiveTransactionId`: References the currently active transaction (null if none active)
- `TransactionSequence`: Counter that increments with each new transaction
- `TotalTransactions`: Total number of transactions created for this customer
- `TotalCompletedTxn`: Count of successfully completed transactions

#### QueueTransaction
Added fields for lifecycle management:
- `SequenceNumber`: Unique sequence per customer (1, 2, 3, ...)
- `DepartmentId`: Direct reference to primary department
- `Status`: Transaction state (Pending → Active → Completed/Cancelled/Skipped)
- `ActivatedAt`: Timestamp when transaction became active
- `CompletedAt`: Timestamp when all items were completed

#### QueueItem
No changes needed - TeamId already provides department context (Team -> Department relationship)

### 2. Transaction Lifecycle States

```
PENDING (0)   → Transaction created, waiting for previous to complete
   ↓
ACTIVE (1)    → Customer called, processing queue items
   ↓
COMPLETED (2) → All queue items done, transaction finished
CANCELLED (3) → Transaction was cancelled by system/admin
SKIPPED (4)   → Customer no-show or skipped
```

### 3. Business Rules

#### Rule 1: Transaction Creation
When creating a new transaction:
```
1. Increment customer.TransactionSequence
2. Assign SequenceNumber = customer.TransactionSequence
3. Set Status = PENDING
4. Set DepartmentId by looking up first QueueItem's Team -> DepartmentId
5. Increment customer.TotalTransactions
```

#### Rule 2: Transaction Activation (Calling Customer)
Before calling a customer's next transaction:
```
1. Check: customer.ActiveTransactionId must be NULL or empty
2. Check: transaction.Status must be PENDING
3. Check: transaction.SequenceNumber == expected next sequence
4. If all checks pass:
   - Set transaction.Status = ACTIVE
   - Set transaction.ActivatedAt = now()
   - Set customer.ActiveTransactionId = transaction.ID
```

#### Rule 3: Transaction Completion
When all queue items in a transaction are done:
```
1. Set transaction.Status = COMPLETED
2. Set transaction.CompletedAt = now()
3. Set customer.ActiveTransactionId = NULL
4. Increment customer.TotalCompletedTxn
5. Trigger event: "transaction.completed" (via RabbitMQ)
```

### 4. MongoDB Indexes (Recommended)

For optimal query performance:

```javascript
// QueueCustomer Collection
db.queueCustomer.createIndex({ "activeTransactionId": 1 })
db.queueCustomer.createIndex({ "queueString": 1 }, { unique: true })

// QueueTransaction Collection
db.queueTransaction.createIndex({ "queueCustomerId": 1, "sequenceNumber": 1 })
db.queueTransaction.createIndex({ "queueCustomerId": 1, "status": 1 })
db.queueTransaction.createIndex({ "departmentId": 1, "status": 1 })
db.queueTransaction.createIndex({ "status": 1, "createdAt": 1 })

// Compound index for finding next transaction to process
db.queueTransaction.createIndex({ 
  "queueCustomerId": 1, 
  "status": 1, 
  "sequenceNumber": 1 
})
```

### 5. Example Workflow

#### Scenario: Student goes to Finance, then Registrar

**Step 1: Student arrives at Finance**
```
Customer: { id: C1, queueString: "2025-001", activeTransactionId: null, transactionSequence: 0 }
```

**Step 2: Create Transaction for Finance**
```
Transaction T1: {
  id: T1,
  queueCustomerId: C1,
  sequenceNumber: 1,
  departmentId: D1 (Finance),
  status: PENDING,
  details: [
    { queueItemId: Q1 (Tuition Payment), order: 1, isCalled: false, isDone: false }
  ]
}

Customer updated: { transactionSequence: 1, totalTransactions: 1 }
```

**Step 3: Finance calls customer**
```
Transaction T1: { status: ACTIVE, activatedAt: timestamp }
Customer updated: { activeTransactionId: T1 }
```

**Step 4: While at Finance window, student also queues for Registrar**
```
Transaction T2: {
  id: T2,
  queueCustomerId: C1,
  sequenceNumber: 2,
  departmentId: D2 (Registrar),
  status: PENDING,
  details: [
    { queueItemId: Q2 (Transcript Request), order: 1, isCalled: false, isDone: false }
  ]
}

Customer updated: { transactionSequence: 2, totalTransactions: 2 }
```

**Step 5: Registrar tries to call customer (BLOCKED)**
```
Check T2.CanBeActivated(Customer):
  - Customer.activeTransactionId = T1 (NOT NULL)
  - T2.ID != T1
  - Result: FALSE ❌

Display: "Customer 2025-001 is currently being served at Finance"
```

**Step 6: Finance completes transaction**
```
Transaction T1: { status: COMPLETED, completedAt: timestamp }
Customer updated: { 
  activeTransactionId: null, 
  totalCompletedTxn: 1 
}

Event published: { type: "transaction.completed", customerId: C1, transactionId: T1 }
```

**Step 7: Now Registrar can call customer**
```
Check T2.CanBeActivated(Customer):
  - Customer.activeTransactionId = null ✅
  - T2.status = PENDING ✅
  - Result: TRUE

Transaction T2: { status: ACTIVE, activatedAt: timestamp }
Customer updated: { activeTransactionId: T2 }
```

### 6. Query Examples

#### Find next transaction to process for a customer
```go
// MongoDB query
filter := bson.M{
    "queueCustomerId": customerId,
    "status": TransactionStatusPending,
}
sort := bson.D{{"sequenceNumber", 1}}
// Returns the pending transaction with lowest sequence number
```

#### Find all customers with active transactions in a department
```go
filter := bson.M{
    "departmentId": departmentId,
    "status": TransactionStatusActive,
}
```

#### Check if customer can receive a new transaction
```go
// A customer can receive new transaction if:
// 1. No active transaction, OR
// 2. Has active transaction but wants to queue for different department
// The new transaction will just be PENDING until active one completes
```

### 7. Event-Driven Updates (via RabbitMQ)

When mutations occur, publish events:

**Events to publish:**
- `transaction.created` - New transaction added
- `transaction.activated` - Customer called, transaction became active
- `transaction.completed` - All items done
- `transaction.cancelled` - Transaction cancelled
- `transaction.item.called` - Specific queue item called
- `transaction.item.completed` - Specific queue item done

**Services should subscribe to:**
- **Display Boards**: Listen to `transaction.activated`, `transaction.completed` to update UI
- **Analytics Service**: Listen to all events for metrics
- **Notification Service**: Listen to `transaction.activated` to send SMS/notifications

### 8. Error Scenarios

#### Scenario: Customer abandons queue (no-show)
```
Action: Set transaction.Status = SKIPPED
Effect: Customer.activeTransactionId = null, allowing next transaction
```

#### Scenario: System detects stale active transaction (>30min)
```
Action: Auto-skip via background job
- Find transactions with Status=ACTIVE and activatedAt < 30min ago
- Set Status = SKIPPED
- Clear customer.activeTransactionId
```

#### Scenario: Customer has emergency and needs to cancel all pending transactions
```
Action: Bulk cancel
- Find all transactions with Status=PENDING for customer
- Set Status = CANCELLED
- Only active transaction remains (if any)
```

## Implementation Checklist

- [x] Update schema models
- [x] Add helper methods for validation
- [ ] Create MongoDB indexes
- [ ] Implement transaction state machine in business logic
- [ ] Add gRPC methods for transaction lifecycle
- [ ] Implement RabbitMQ event publishers
- [ ] Add background job for stale transaction cleanup
- [ ] Update GraphQL resolvers to check transaction constraints
- [ ] Add unit tests for helper methods
- [ ] Add integration tests for cross-department scenarios
