package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rishabh-iitj2029/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionstring = "mongoDBURI"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	// 1. Configure the options
	clientOption := options.Client().ApplyURI(connectionstring)

	// 2. Initialize the client wrapper (does not test network)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Configuration error: ", err)
	}

	// 3. MANDATORY: Actually verify the network connection
	// We use a timeout context so your app doesn't hang forever if MongoDB is down
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	fmt.Println("MongoDB connection done")

	// 4. Get the collection reference
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

	collection = client.Database(dbName).Collection(colName)

	//collection instance

	fmt.Println("Collection instance is ready")
}



func insertOneMovie(movie models.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println("Inserted the one movie", inserted)
}


