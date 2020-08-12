package main

import (
	"fmt"
	"garagesvc/module/mongodb"
	"garagesvc/route"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	mongodb.Connect()
}

func main() {
	fmt.Println(time.Now())

	//Echo
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	//Add routers
	route.Bootstrap(e)

	// Load dotenv for port option
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("PORT"))
	//Open port
	e.Start(os.Getenv("PORT"))
}
