package controller

import (
	"employeeregister/database"
	"employeeregister/models"
	"employeeregister/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var jwtkey = []byte(os.Getenv("JWT_KEY"))
var logger = logrus.New()

func init() {
	// Set up logging configuration
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
}

func GenerateJwt(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}

func UserRegister(c *gin.Context) {
	var UserRegister models.User

	if err := c.BindJSON(&UserRegister); err != nil {
		logger.Errorf("Error binding JSON: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateEmail(UserRegister.Email); err != nil {

		logger.Errorf("Invalid email fomar", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return

	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", UserRegister.Email).First(&existingUser).Error; err == nil {
		logger.Errorf("Email already registered: %v", UserRegister.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
		return
	}

	if err := utils.Validatepassword(strings.TrimSpace(UserRegister.Password)); err != nil {

		logger.Errorf("Error hashing password:%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserRegister.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Errorf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	UserRegister.Password = string(hashedPassword)
	var ExistingPassowrd models.User
	if err := database.DB.Where("password = ?", string(hashedPassword)).First(&ExistingPassowrd).Error; err == nil {

		logger.Errorf("Password is already taken by  user: %v", UserRegister.Password)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is already taken by  user"})
		return
	}
	result := database.DB.Create(&UserRegister)
	if result.Error != nil {
		logger.Errorf("Error creating user: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	logger.Infof("User registered successfully: %v", UserRegister)
	c.JSON(http.StatusOK, &UserRegister)
}

func UserGet(c *gin.Context) {
	var UserGet []models.User

	result := database.DB.Find(&UserGet)
	if result.Error != nil {
		logger.Errorf("Error retrieving users: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(UserGet) == 0 {
		logger.Warn("No users found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data found"})
		return
	}

	logger.Infof("Retrieved users: %v", UserGet)
	c.JSON(http.StatusOK, &UserGet)
}

func UseDelete(c *gin.Context) {
	param := c.Param("id")
	var userData models.User

	result := database.DB.First(&userData, param)
	if result.Error != nil {
		logger.Errorf("Error finding user: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	database.DB.Delete(&userData, param)
	logger.Infof("User deleted successfully: %s", param)
	c.JSON(http.StatusOK, "User delete successful")
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	body := models.User{}
	var updateUser models.User

	result := database.DB.First(&updateUser, id)
	if result.Error != nil {
		logger.Errorf("Error finding user for update: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		logger.Errorf("Error binding JSON for update: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateUser.FirstName = body.FirstName
	updateUser.LastName = body.LastName
	updateUser.Email = body.Email
	updateUser.Password = body.Password

	database.DB.Save(&updateUser)
	logger.Infof("User updated successfully: %v", updateUser)
	c.JSON(http.StatusOK, &updateUser)
}

func UserGetbyEmail(c *gin.Context) {
	var user []models.User
	id := c.Param("id")

	getbyId := database.DB.First(&user, id)
	if getbyId.Error != nil {
		logger.Errorf("Error finding user by email: %v", getbyId.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	logger.Infof("User retrieved by email: %v", user)
	c.JSON(http.StatusOK, &user)
}

func UserLogin(c *gin.Context) {
	var user models.User
	var logindata struct {
		Email    string
		Password string
	}

	if err := c.BindJSON(&logindata); err != nil {
		logger.Errorf("Error binding JSON for login: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := database.DB.Where("email=?", logindata.Email).First(&user)
	if result.Error != nil {
		logger.Errorf("Error finding user for login: %v", result.Error)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logindata.Password)); err != nil {
		logger.Errorf("Password mismatch for user: %v", logindata.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := GenerateJwt(user.Email)
	if err != nil {
		logger.Errorf("Error generating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	logger.Infof("User logged in successfully: %v", user.Email)
	c.JSON(http.StatusOK, &token)
}
