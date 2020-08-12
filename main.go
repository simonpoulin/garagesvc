package main

import (
	"fmt"
	"garagesvc/module/mongodb"
	"garagesvc/route"
	"os"
	"time"

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

	//Open port
	e.Start(":" + os.Getenv("PORT"))
}
