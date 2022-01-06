package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CatsCollection *mongo.Collection
	// Declare Context type object for managing multiple API requests
	Ctx, _ = context.WithTimeout(context.Background(), 15*time.Second)
)

// Setup opens a database connection to mongodb
func Setup() error {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := fmt.Sprintf("mongodb://%s:%s/", host, port)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		return err
	}

	db := client.Database("api")
	CatsCollection = db.Collection("cats")

	return nil
}
