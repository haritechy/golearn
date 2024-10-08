package main

import (
	"employeeregister/database"
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

	routes.EmployeeRoutes(r)
	routes.WarrantyRoutes(r)
	routes.UserRoutes(r)

	r.Run(":8080")
}
