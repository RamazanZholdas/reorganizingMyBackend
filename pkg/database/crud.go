package database

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateCollection(collectionName string) error {
	err := m.Db.CreateCollection(context.TODO(), collectionName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return fmt.Errorf("%s collection already exists", collectionName)
		} else {
			return err
		}
	}
	return nil
}

func (m *MongoDB) FindOne(collectionName string, filter interface{}, result interface{}) error {
	err := m.Db.Collection(collectionName).FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) FindMany(collectionName string, filter interface{}, results interface{}) error {
	cur, err := m.Db.Collection(collectionName).Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var elem interface{}
		if err := cur.Decode(&elem); err != nil {
			return err
		}
		results = append(results.([]interface{}), elem)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateOne(collectionName string, filter interface{}, update interface{}) error {
	_, err := m.Db.Collection(collectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteOne(collectionName string, filter interface{}) error {
	_, err := m.Db.Collection(collectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) InsertOne(collectionName string, data interface{}) error {
	collection := m.Db.Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) InsertMany(collectionName string, data []interface{}) error {
	collection := m.Db.Collection(collectionName)
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) CountProducts(collectionName string) (int64, error) {
	count, err := m.Db.Collection(collectionName).CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}
