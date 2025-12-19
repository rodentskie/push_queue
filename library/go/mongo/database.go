package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DatabaseInfo) GetCollection(mongoClient *mongo.Client, col string) *mongo.Collection {
	readOnlyDb := mongoClient.Database(db.Database)

	readOnlyCol := readOnlyDb.Collection(col)

	return readOnlyCol
}
