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

	fmt.Println("Connected to MongoDB")
	DB = client
	return client
}



//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("blog-golang").Collection(collectionName)
	return collection
}
