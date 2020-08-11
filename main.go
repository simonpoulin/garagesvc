package main

import (
	"fmt"
	controller "garagesvc/controllers"
	validator "garagesvc/validators"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println(time.Now())
	//Echo
	e := echo.New()

	e.Use(middleware.CORS())

	//APIs
	//Company
	// e.GET("/companies", GetCompanies)
	// e.GET("/companies/:id", GetCompany)
	// e.POST("/companies", CreateCompany)
	// e.PATCH("/companies/:id", UpdateCompany)
	// e.DELETE("/companies/:id", DeleteCompany)

	//Service
	// e.GET("/services", GetServices)
	// e.GET("/services/:id", GetService)
	// e.POST("/services", CreateService)
	// e.PATCH("/services/:id", UpdateService)
	// e.DELETE("/services/:id", DeleteService)

	//User
	// e.GET("/users", GetUsers)
	// e.GET("/users/:id", GetUser)
	// e.POST("/users", CreateUser)
	// e.PATCH("/users/:id", UpdateUser)
	// e.DELETE("/users/:id", DeleteUser)

	//Employee
	// e.GET("/employees", GetEmployees)
	// e.GET("/employees/:id", GetEmployee)
	e.POST("/employees", controller.EmployeeCreate, validator.EmployeeCreate)
	e.POST("/employees/login", controller.EmployeeLogin, validator.EmployeeLogin)
	// e.PATCH("/employees/:id", UpdateEmployee)
	// e.DELETE("/employees/:id", DeleteEmployee)

	//Book
	// e.GET("/books", GetBooks)
	// e.GET("/books/:id", GetBook)
	// e.GET("/books/status/:status", GetBooksByStatus)
	// e.POST("/books", CreateBook)
	// e.PATCH("/books/:id", UpdateBook)
	// e.DELETE("/books/:id", DeleteBook)

	e.Start(":9998")
}
