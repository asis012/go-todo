package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "go_todo"

// Collection name
const collName = "todolist"

// collection object/instance
var Collection *mongo.Collection

// create connection with mongo db
func SetupDB() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	Collection = client.Database(dbName).Collection(collName)
}
