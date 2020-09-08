package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env ENV

// ENV ...
type ENV struct {
	Port               string
	CustomerKey        string
	EmployeeKey        string
	DatabaseConnection string
	DatabaseName       string
}

// Init ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	env.Port = os.Getenv("PORT")
	env.CustomerKey = os.Getenv("CUSTOMER_KEY")
	env.EmployeeKey = os.Getenv("EMPLOYEE_KEY")
	env.DatabaseConnection = os.Getenv("DATABASE_CONNECTION")
	env.DatabaseName = os.Getenv("DATABASE_NAME")
}

// GetENV ...
func GetENV() ENV {
	return env
}
