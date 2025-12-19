package cursor

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCursorGenerate(t *testing.T) {
	objectID := primitive.NewObjectID()
	createdAt := time.Now().Unix()

	cursor := CursorGenerate(createdAt, objectID[:])

	if len(cursor) != 20 {
		t.Errorf("Expected cursor length 20, got %d", len(cursor))
	}

	// Verify timestamp is correctly stored in first 8 bytes
	storedTimestamp := int64(binary.BigEndian.Uint64(cursor[:8]))
	if storedTimestamp != createdAt {
		t.Errorf("Expected timestamp %d, got %d", createdAt, storedTimestamp)
	}

	// Verify ObjectID is correctly stored in last 12 bytes
	if !bytes.Equal(cursor[8:], objectID[:]) {
		t.Error("Expected ObjectID to match stored bytes")
	}
}

func TestCursorGenerateWithZeroTimestamp(t *testing.T) {
	objectID := primitive.NewObjectID()
	createdAt := int64(0)

	cursor := CursorGenerate(createdAt, objectID[:])

	if len(cursor) != 20 {
		t.Errorf("Expected cursor length 20, got %d", len(cursor))
	}

	storedTimestamp := int64(binary.BigEndian.Uint64(cursor[:8]))
	if storedTimestamp != createdAt {
		t.Errorf("Expected timestamp %d, got %d", createdAt, storedTimestamp)
	}
}

func TestCursorGenerateOrdering(t *testing.T) {
	objectID1 := primitive.NewObjectID()
	objectID2 := primitive.NewObjectID()

	// Earlier timestamp
	time1 := time.Now().Unix()
	time.Sleep(1 * time.Millisecond)
	// Later timestamp
	time2 := time.Now().Unix()

	cursor1 := CursorGenerate(time1, objectID1[:])
	cursor2 := CursorGenerate(time2, objectID2[:])

	// cursor1 should be "less than" cursor2 when compared bytewise
	if bytes.Compare(cursor1, cursor2) >= 0 {
		t.Error("Expected cursor1 < cursor2 based on timestamp ordering")
	}
}
