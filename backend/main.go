package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var employeeCollection *mongo.Collection

func main() {
	// MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	employeeCollection = client.Database("employeeDB").Collection("employees")

	// Set up Gin router
	router := gin.Default()

	// API Routes
	router.GET("/employees", getEmployees)
	router.GET("/employees/:id", getEmployeeByID)
	router.POST("/employees", addEmployee)
	router.PUT("/employees/:id", updateEmployee)
	router.DELETE("/employees/:id", deleteEmployee)

	// Start server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
