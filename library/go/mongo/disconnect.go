package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DatabaseInfo) Disconnect(ctx context.Context, mongoClient *mongo.Client) error {

	if err := mongoClient.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}
