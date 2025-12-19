package redis

import (
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	client := New("localhost:6379", "", 0)
	if client == nil {
		t.Fatal("Expected non-nil redis client")
	}
}

func TestPutIDAndGetID(t *testing.T) {
	ctx := context.Background()
	client := New("localhost:6379", "", 0)
	sheet := "test_sheet"
	originalID := "test_id"
	generated := []byte("test_value")
	ttl := 10 * time.Second

	err := PutID(ctx, client, sheet, originalID, generated, ttl)
	if err != nil {
		t.Fatalf("PutID failed: %v", err)
	}

	val, err := GetID(ctx, client, sheet, originalID)
	if err != nil {
		t.Fatalf("GetID failed: %v", err)
	}
	if string(val) != string(generated) {
		t.Errorf("Expected %s, got %s", generated, val)
	}
}
