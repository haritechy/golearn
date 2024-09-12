package models

import (
	"time"

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

	FirstName string
	LastName  string
	Password  string
	Email     string
}

type Claims struct {
	Email string `json:"email"`

	jwt.StandardClaims
}

type WrrantyData struct {
	ID              uint      `gorm:"primaryKey"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       *time.Time
	Vendor          string
	WarrantyID      string
	Name            string
	MonthlyPrice    string
	Discount        string
	AnnualPrice     string
	PlanDescription string
	Status          string
	Picture         string
	PictureName     string
}
