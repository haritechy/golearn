package controller

import (
	"employeeregister/database"
	"employeeregister/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		logger.WithError(err).Error("error Binding Json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Create(&employee)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &employee)
}

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	result := database.DB.Find(&employees)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &employees)
}

func GetEmployee(c *gin.Context) {

	var id = c.Param("id")
	var Employee models.Employee

	// database.DB.First(&Employee, id)
	result := database.DB.First(&Employee, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, &Employee)

}
func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	var Employee models.Employee

	result := database.DB.First(&Employee, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	database.DB.Delete(&Employee, id)
	c.JSON(http.StatusOK, "Succefuly deleted")
}

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")
	body := models.Employee{}
	var Employee models.Employee
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result := database.DB.First(&Employee, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	Employee.Name = body.Name
	Employee.Position = body.Position
	Employee.Salary = body.Salary
	Employee.Email = body.Email
	database.DB.Save(&Employee)
	c.JSON(http.StatusOK, &Employee)
}
