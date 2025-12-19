package mongo

import (
	"context"
	"testing"
)

func TestDisconnect(t *testing.T) {

	dbInfo := DatabaseInfo{
		Uri: "mongodb://localhost:27017",
	}

	ctx := context.Background()
	mongoClient, err := dbInfo.Connect(ctx)

	if mongoClient == nil {
		t.Error("Expected a non-nil mongoClient, but got nil")
	}

	if err != nil {
		t.Errorf("Error connection to mongoDb: %v\n", err)
	}

	if err := dbInfo.Disconnect(ctx, mongoClient); err != nil {
		t.Errorf("Expected a non-nil error, but got %v\n", err)
	}
}
