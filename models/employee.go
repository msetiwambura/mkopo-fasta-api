package models

import (
	"time"
)

type Employee struct {
	ID          uint      `gorm:"primaryKey"`
	FirstName   string    `gorm:"size:50;not null"`
	LastName    string    `gorm:"size:50;not null"`
	Gender      string    `gorm:"type:enum('Male', 'Female', 'Other');not null"`
	DateOfBirth time.Time `gorm:"not null"`
	NationalID  string    `gorm:"size:20;unique;not null"`
	Address     string    `gorm:"size:255;not null"`
	City        string    `gorm:"size:100;not null"`
	Province    string    `gorm:"size:100;not null"`
	PostalCode  string    `gorm:"size:10;not null"`
	Country     string    `gorm:"size:100;not null"`
	PhoneNumber string    `gorm:"size:15;not null"`
	Email       string    `gorm:"size:100;unique;not null"`
	JobTitle    string    `gorm:"size:100;not null"`
	Department  string    `gorm:"size:100;not null"`
	Salary      float64   `gorm:"not null"`
	HireDate    time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
