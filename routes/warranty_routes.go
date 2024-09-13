package routes

import (
	"employeeregister/controller"

	"github.com/gin-gonic/gin"
)

func WarrantyRoutes(r *gin.Engine) {
	warrantyRoutes := r.Group("/warranty")
	// warrantyRoutes.Use(
	// 	middleware.AuthorizeJWT())
	{
		warrantyRoutes.POST("/create", controller.WarrntyUpload)
		warrantyRoutes.GET("/all", controller.WaarntyGet)
		warrantyRoutes.PUT("/update/:id", controller.WarrantyUpdate)
		warrantyRoutes.GET("/pending", controller.PendingWarranty)
		warrantyRoutes.DELETE("/delete/:id", controller.DeleteWarranty)
		warrantyRoutes.POST("/upload", controller.ExcelUpload)
		warrantyRoutes.DELETE("/deletes/:id", controller.WrrantyDatagetDelete)
		warrantyRoutes.GET("/alls", controller.Warrantydatget)
		warrantyRoutes.PUT("/ups/:id", controller.WarrntyDataUpdate)

	}
}
