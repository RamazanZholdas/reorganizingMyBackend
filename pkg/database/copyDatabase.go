package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CopyDatabase(client *mongo.Client, oldDBName string, newDBName string) error {
	oldDB := client.Database(oldDBName)
	newDB := client.Database(newDBName)

	collections, err := oldDB.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	for _, collectionName := range collections {
		oldCollection := oldDB.Collection(collectionName)
		newCollection := newDB.Collection(collectionName)

		cursor, err := oldCollection.Find(context.Background(), bson.D{})
		if err != nil {
			return err
		}

		var documents []interface{}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var document interface{}
			err = cursor.Decode(&document)
			if err != nil {
				return err
			}
			documents = append(documents, document)
		}
		if err := cursor.Err(); err != nil {
			return err
		}

		_, err = newCollection.InsertMany(context.Background(), documents)
		if err != nil {
			return err
		}
	}

	return nil
}
