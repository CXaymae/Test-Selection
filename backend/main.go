package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
	"github.com/CXaymae/Test-Selection/backend/models" 
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
	"github.com/go-playground/validator/v10"
    "net/http"
)

var client *mongo.Client
var employeeCollection *mongo.Collection
var validate = validator.New()

func main() {
    // Connect to MongoDB
    ctx := context.TODO()
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Check the MongoDB connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB")

    // Get the Employee collection
    employeeCollection = client.Database("employeeDB").Collection("employees")

    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/employees", getEmployees).Methods("GET")
    r.HandleFunc("/employees/{id}", getEmployeeByIDHandler).Methods("GET")
    r.HandleFunc("/employees", createEmployeeHandler).Methods("POST")
    r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
    r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

    // Start the server
    fmt.Println("Server started on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
    ctx := context.TODO()
    cursor, err := employeeCollection.Find(ctx, bson.D{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var employees []models.Employee
    for cursor.Next(ctx) {
        var employee models.Employee
        if err := cursor.Decode(&employee); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        employees = append(employees, employee)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employees)
}

// Get employee by ID
func getEmployeeByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    employeeID := vars["id"]

    ctx := context.TODO()
    var employee models.Employee
    err := employeeCollection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&employee)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Employee not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employee)
}

// Create a new employee
func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
    var employee models.Employee
    err := json.NewDecoder(r.Body).Decode(&employee)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Perform validation
    if err := validate.Struct(employee); err != nil {
        http.Error(w, fmt.Sprintf("Validation failed: %s", err.Error()), http.StatusBadRequest)
        return
    }

    // Insert the employee into MongoDB
    ctx := context.TODO()
    employee.DateOfHire = time.Now() // Set the DateOfHire to the current time
    result, err := employeeCollection.InsertOne(ctx, employee)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

// Update an existing employee
func updateEmployee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    employeeID := vars["id"]

    var updatedEmployee models.Employee
    err := json.NewDecoder(r.Body).Decode(&updatedEmployee)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    ctx := context.TODO()
    update := bson.M{
        "$set": updatedEmployee,
    }

    result, err := employeeCollection.UpdateOne(ctx, bson.M{"_id": employeeID}, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result.MatchedCount == 0 {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedEmployee)
}

// Delete an employee by ID
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    employeeID := vars["id"]

    ctx := context.TODO()
    result, err := employeeCollection.DeleteOne(ctx, bson.M{"_id": employeeID})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Employee deleted")
}

// writeErrorResponse is a utility function that formats an error response 
// and sends it to the client. It sets the appropriate HTTP status code and 
// a JSON response with an error message.

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    // Set the Content-Type header to 'application/json', indicating that
    // the response will be in JSON format.
    w.Header().Set("Content-Type", "application/json")
    
    // Set the HTTP status code for the response. This reflects the type of error
    // that occurred (e.g., 404 for not found, 500 for server error).
    w.WriteHeader(statusCode)
    
    // Encode the error message into a JSON object and send it in the response body.
    // The map here is used to structure the JSON object with a single "error" key.
    json.NewEncoder(w).Encode(map[string]string{"error": message})
}
