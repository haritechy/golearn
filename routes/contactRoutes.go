package routes

import (
	"employeeregister/controller"

	"github.com/gin-gonic/gin"
)

func ContactRoutes(r *gin.Engine) {

	userMessageRoutes := r.Group("/message")

	{

		userMessageRoutes.POST("/contact", controller.UserMessage)
	}

}
