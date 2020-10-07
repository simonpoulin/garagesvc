package mongodb

import (
	"context"
	"garagesvc/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// Connect ...
func Connect() {

	// Load dotenv for database connection
	env := config.GetENV()

	// Get connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(env.DatabaseConnection))
	if err != nil {
		log.Fatal(err)
	}

	// Set data
	db = client.Database(env.DatabaseName)
}

// BookingCol ...
func BookingCol() *mongo.Collection {
	return db.Collection("Bookings")
}

// CompanyCol ...
func CompanyCol() *mongo.Collection {
	return db.Collection("Companies")
}

// EmployeeCol ...
func EmployeeCol() *mongo.Collection {
	return db.Collection("Employees")
}

// ServiceCol ...
func ServiceCol() *mongo.Collection {
	return db.Collection("Services")
}

// CustomerCol ...
func CustomerCol() *mongo.Collection {
	return db.Collection("Customers")
}

// ResourceCol ...
func ResourceCol() *mongo.Collection {
	return db.Collection("Resources")
}
