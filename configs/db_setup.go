package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client 
var dbName string

func ConnectDB() *mongo.Client {

	mongoUri := MongoUri()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer cancelContext()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbName = "blog-golang"
	fmt.Println("Connected to MongoDB")
	DB = client
	return client
}



//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}


func DropCollection(client *mongo.Client, collectionName string) error {
	collection := client.Database(dbName).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	return collection.Drop(ctx)
}


func ConnectDBTest() *mongo.Client {

	mongoUri := MongoUriTest()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelContext := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer cancelContext()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbName = "blog-golang-test"
	fmt.Println("Connected to MongoDB tesing")
	DB = client
	return client
}

