package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name     string  `json:"name"  binding:"required" gorm:"not null"`
	Email    string  `json:"email"  binding:"required" gorm:"not null"`
	Position string  `json:"Position"  binding:"required" gorm:"not null"`
	Salary   float64 `json:"Salary"  binding:"required" gorm:"not null"`
}

type Warranty struct {
	gorm.Model

	ProductName  string `json:"productname"  binding:"required" gorm:"not null"`
	WarrantyId   int
	Status       string `json:"status" binding:"required" gorm:"not null"`
	Vendor       string `json:"vendor"  binding:"required" gorm:"not null"`
	MonthlyPrice int    `json:"monthlyprice"  binding:"required" gorm:"not null"`
	Discount     int    `json:"discount"  binding:"required" gorm:"not null"`
	AnnualPrice  int    `json:"annualPrice"  binding:"required" gorm:"not null"`
}
type User struct {
	gorm.Model

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname" `
	FullName  string `json:"fullname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Otp       string `json:"otp"`
}

type Claims struct {
	Email string `json:"email"  binding:"required" gorm:"not null"`

	jwt.StandardClaims
}

type WrrantyData struct {
	ID              uint      `gorm:"primaryKey"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       *time.Time
	Vendor          string `json:"vendor"  binding:"required" gorm:"not null"`
	WarrantyID      string `json:"warrantyid"`
	Name            string `json:"name"  binding:"required" gorm:"not null"`
	MonthlyPrice    string `json:"monthlyprice"  binding:"required" gorm:"not null"`
	Discount        string `json:"discount" `
	AnnualPrice     string `json:"annualprice"  binding:"required" gorm:"not null"`
	PlanDescription string `json:"plandescription"  binding:"required" gorm:"not null"`
	Status          string `json:"status"  binding:"required" gorm:"not null"`
	Picture         string `json:"picture"  binding:"required" gorm:"not null"`
	PictureName     string `json:"picturename"   binding:"required" gorm:"not null"`
}

type Votes struct {
	ID          uint `gorm:"primaryKey"`
	VoterId     int  `json:"voteId"`
	CandidateId int  `json:"candidateid"`
}
type Candidate struct {
	Id        uint   `gorm:"PrimaryKey"`
	Name      string `gorm:"name"`
	VoteCount int    `gorm:"default:0"`
}

type UserMessage struct {
	Name        string `json:"name"`
	Email       string `json:"email" binding:"required"`
	Message     string `json:"message" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
}
type Otp struct {
	Otp string
}
