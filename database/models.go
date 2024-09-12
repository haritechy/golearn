package database

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name     string
	Email    string
	Position string
	Salary   float64
}

type Warranty struct {
	gorm.Model

	ProductName  string
	WarrantyId   int
	Status       string
	Vendor       string
	MonthlyPrice int
	Discount     int
	AnnualPrice  int
}
type User struct {
	gorm.Model

	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Password  string `binding:"required"`
	Email     string `binding:"required"`
}

type Claims struct {
	Email string `json:"email"`

	jwt.StandardClaims
}
