package main

import (
	"garagesvc/module/mongodb"
	"garagesvc/route"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	mongodb.Connect()
}

func main() {
	//Load dotenv for port option
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//Echo
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	//Add routers
	route.Bootstrap(e)

	//Open port
	e.Start(os.Getenv("PORT"))
}
