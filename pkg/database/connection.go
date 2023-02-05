package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func (m *MongoDB) Connect(uri, database string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error creating new MongoDB client: %v", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	m.Client = client
	m.Db = client.Database(database)

	err = m.Db.RunCommand(context.TODO(), bson.D{{Key: "create", Value: database}}).Err()
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("warning: database already exists")
		} else {
			return fmt.Errorf("error creating database: %v", err)
		}
	}

	return nil
}

func (m *MongoDB) Disconnect() {
	err := m.Client.Disconnect(context.TODO())
	if err != nil {
		log.Println("Error disconnecting from MongoDB:", err)
	}
}
