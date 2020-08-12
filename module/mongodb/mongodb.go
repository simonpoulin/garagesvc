package mongodb

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

// Connect ...
func Connect() {
	// Load dotenv for database connection
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// Get connection
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DATABASE_CONNECTION")))
	if err != nil {
		log.Fatal(err)
	}

	// Set data
	db = client.Database(os.Getenv("DATABASE_NAME"))
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

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection("Users")
}
