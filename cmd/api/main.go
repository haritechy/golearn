package main

import (
	"employeeregister/controller"
	"employeeregister/database"
	"employeeregister/middleware"
	"employeeregister/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {

		log.Fatal("error loding .env")
	}
	gin.SetMode(gin.ReleaseMode)
	database.Connect()
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	routes.EmployeeRoutes(r)
	routes.WarrantyRoutes(r)
	routes.UserRoutes(r)
	routes.ContactRoutes(r)
	routes.VoterRoutes(r)
	controller.Arrays()
	r.Run(":8080")

}
