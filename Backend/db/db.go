package db

import (
	"context"
	"fmt"

	"github.com/shaply/ProximityChat/Backend/config"
	"github.com/shaply/ProximityChat/Backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI string
}

func NewMongoDBStorage(cfg Config) (*mongo.Database, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.URI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Pinged your deployment. Successfully connected to MongoDB Cluster!")

	// Get the database for ProximityChat
	db := client.Database(types.ProximityChat.Database)

	return db, nil
}

// Create the connection with the database
func InitiateConnection() (*mongo.Database, error) {
	return NewMongoDBStorage(Config{
		URI: config.Envs.MongoDB_URI,
	})
}
