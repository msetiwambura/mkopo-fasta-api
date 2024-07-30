package models

import (
	"time"
)

type Customer struct {
	ID               uint   `gorm:"primaryKey"`
	FirstName        string `gorm:"size:50;not null"`
	LastName         string `gorm:"size:50;not null"`
	Gender           string `gorm:"type:enum('Male', 'Female', 'Other');not null"`
	DateOfBirth      string `json:"dateOfBirth" gorm:"type:date;not null"`
	NationalID       string `gorm:"size:20;unique;not null"`
	Address          string `gorm:"size:255;not null"`
	City             string `gorm:"size:100;not null"`
	Province         string `gorm:"size:100;not null"`
	PostalCode       string `gorm:"size:10;not null"`
	Country          string `gorm:"size:100;not null"`
	PhoneNumber      string `gorm:"size:15;not null"`
	Email            string `gorm:"size:100;unique;not null"`
	EmploymentStatus string `gorm:"type:enum('Employed', 'Unemployed', 'Self-employed', 'Student', 'Retired');not null"`
	AnnualIncome     float64
	CreditScore      int
	CreatedAt        time.Time         `gorm:"autoCreateTime"`
	UpdatedAt        time.Time         `gorm:"autoUpdateTime"`
	Loans            []Loan            `gorm:"foreignKey:CustomerID"`
	LoanApplications []LoanApplication `gorm:"foreignKey:CustomerID"`
	Transactions     []Transaction     `gorm:"foreignKey:CustomerID"`
}

type CustomerRes struct {
	ID               uint      `json:"Id"`
	FirstName        string    `json:"FirstName"`
	LastName         string    `json:"LastName"`
	Gender           string    `json:"Gender"`
	DateOfBirth      string    `json:"DateOfBirth"`
	NationalID       string    `json:"NationalID"`
	Address          string    `json:"Address"`
	City             string    `json:"City"`
	Province         string    `json:"Province"`
	PostalCode       string    `json:"PostalCode"`
	Country          string    `json:"Country"`
	PhoneNumber      string    `json:"PhoneNumber"`
	Email            string    `json:"Email"`
	EmploymentStatus string    `json:"EmploymentStatus"`
	AnnualIncome     float64   `json:"AnnualIncome"`
	CreditScore      int       `json:"CreditScore"`
	CreatedAt        time.Time `json:"CreatedAt"`
	UpdatedAt        time.Time `json:"UpdatedAt"`
}
