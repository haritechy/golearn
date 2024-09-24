package controller

import (
	"employeeregister/cmd/api/helpers"
	"employeeregister/database"
	"employeeregister/models"
	"employeeregister/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UserMessage(c *gin.Context) {

	var UserMessage models.UserMessage

	if err := c.BindJSON(&UserMessage); err != nil {

		logger.WithError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return

	}
	if emailValidation := utils.ValidateEmail(UserMessage.Email); emailValidation != nil {

		logger.Errorf("please enetr crt valid email address")
		c.JSON(http.StatusBadRequest, gin.H{"error": "please enter valid email address"})
		return
	}

	database.DB.Create(&UserMessage)
	userWelcome, userEmailBody, adminEmailBody, sheduled := helpers.Stringshelp(UserMessage.Name, UserMessage.Message, UserMessage.Email, UserMessage.PhoneNumber)
	if err :=
		utils.SendEmail(UserMessage.Email, userWelcome, userEmailBody); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not send to email to user"})
		return
	}
	if err :=
		utils.SendEmail(os.Getenv("MAIL"), sheduled, adminEmailBody); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not send to email to admin"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message received and email sento to both user and admin"})
}
