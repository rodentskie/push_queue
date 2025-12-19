package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (db *DatabaseInfo) Connect(ctx context.Context) (*mongo.Client, error) {

	mongoconn := options.Client().ApplyURI(db.Uri)
	mongoClient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {

		return nil, err

	}

	return mongoClient, nil
}
