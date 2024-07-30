package models

import (
	"time"
)

type Guarantor struct {
	ID          uint `gorm:"primaryKey"`
	LoanID      uint
	Loan        Loan
	FirstName   string    `gorm:"size:50;not null"`
	LastName    string    `gorm:"size:50;not null"`
	NationalID  string    `gorm:"size:20;unique;not null"`
	Address     string    `gorm:"size:255;not null"`
	City        string    `gorm:"size:100;not null"`
	Province    string    `gorm:"size:100;not null"`
	PostalCode  string    `gorm:"size:10;not null"`
	Country     string    `gorm:"size:100;not null"`
	PhoneNumber string    `gorm:"size:15;not null"`
	Email       string    `gorm:"size:100;unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
