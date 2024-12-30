package db

import (
	"context"
	"log"
	"time"

	"github.com/shubhamrwtktw/lic/internal/config/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client     *mongo.Client
	database   *mongo.Database
	Collection Collection
}

type Collection struct {
	Users *mongo.Collection
}

var DB *Database

func Connect() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(env.Env.MONGODB_URL))

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}


	log.Println("Connected from MongoDB")

	db := client.Database("lic")

	DB = &Database{
		client:   client,
		database: db,
		Collection: Collection{
			Users: db.Collection("users"),
		},
	}

}

func CloseDB() {
	if DB.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		DB.client.Disconnect(ctx)
		log.Println("Disconnected from MongoDB")
	}
}
