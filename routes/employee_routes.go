package routes

import (
	"employeeregister/controller"
	"employeeregister/middleware"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	employeeRoutes := r.Group("/employees")
	employeeRoutes.Use(middleware.AuthorizeJWT())
	{
		employeeRoutes.GET("/", controller.GetEmployees)
		employeeRoutes.POST("/create", controller.CreateEmployee)
		employeeRoutes.DELETE("/:id", controller.DeleteEmployee)
		employeeRoutes.GET("/:id", controller.GetEmployee)
		employeeRoutes.PUT("/update/:id", controller.UpdateEmployee)
	}
}
