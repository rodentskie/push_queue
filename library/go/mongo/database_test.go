package mongo

import (
	"context"
	"testing"
)

func TestGetCollection(t *testing.T) {

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

	readOnly := dbInfo.GetCollection(mongoClient, "readOnly")

	if readOnly == nil {
		t.Error("Expected a non-nil readOnly, but got nil")
	}

}
