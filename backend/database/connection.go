package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variable to hold the MongoDB database instance
var DB *mongo.Database

// ConnectDB initializes the connection to the MongoDB database
func ConnectDB() {
	// Create a new MongoDB client with the provided connection URI
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		// Log and terminate the application if the client creation fails
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Set up a context with a timeout to avoid hanging connection attempts
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure resources are released when the context expires

	// Connect to the MongoDB server using the created client and context
	err = client.Connect(ctx)
	if err != nil {
		// Log and terminate the application if the connection fails
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Select the database to use and assign it to the global variable
	DB = client.Database("employee_management")
	log.Println("Connected to MongoDB")

	// Ensure the client disconnects properly when the application exits
	defer func() {
		// Attempt to disconnect the client
		if err := client.Disconnect(ctx); err != nil {
			// Log any errors that occur during disconnection
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()
}
