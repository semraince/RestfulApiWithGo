package database

import (
	"GoDatabaseAssignment/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var db *mongo.Database

func Init() error {
	dbUri := config.Get().Servers.DB.Uri
	clientOptions := options.Client()
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(20)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri), clientOptions)
	if err != nil {
		return err
	}
	db = client.Database(config.Get().Servers.DB.Name)
	return nil
}

func Aggregate(collection string, filter interface{}, results interface{}) error {
	cursor, err := db.Collection(collection).Aggregate(context.TODO(), filter)
	if nil != err {
		log.Printf("DB Aggregate error %v\n", err)
		return err
	}

	err = cursor.All(context.TODO(), results)
	if nil != err {
		log.Printf("DB Aggregate error %v\n", err)
		return err
	}
	return nil
}
