package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	env ENV
	ext AcceptedExt
	img DefaultImg
)

// ENV ...
type ENV struct {
	Port               string
	CustomerKey        string
	EmployeeKey        string
	DatabaseConnection string
	DatabaseName       string
	ImageDirectory     string
}

// AcceptedExt ...
type AcceptedExt struct {
	Extensions []string
}

// DefaultImg ...
type DefaultImg struct {
	SmallImage  string
	MediumImage string
	LargeImage  string
}

// Init ...
func Init() {
	// Get dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Set env variables
	env.Port = os.Getenv("PORT")
	env.CustomerKey = os.Getenv("CUSTOMER_KEY")
	env.EmployeeKey = os.Getenv("EMPLOYEE_KEY")
	env.DatabaseConnection = os.Getenv("DATABASE_CONNECTION")
	env.DatabaseName = os.Getenv("DATABASE_NAME")
	env.ImageDirectory = os.Getenv("IMAGE_DIR")

	// Set accepted image extensions
	ext.Extensions = append(ext.Extensions, ".png")
	ext.Extensions = append(ext.Extensions, ".jpg")
	ext.Extensions = append(ext.Extensions, ".jpeg")

	// Set default image
	img.SmallImage = "5f7691bee77eedd2166d4343.png"
	img.MediumImage = "5f7691bee77eedd2166d4342.png"
	img.LargeImage = "5f7691bee77eedd2166d4341.png"
}

// GetENV ...
func GetENV() ENV {
	return env
}

// GetEXT ...
func GetEXT() AcceptedExt {
	return ext
}

// GetIMG ...
func GetIMG() DefaultImg {
	return img
}
