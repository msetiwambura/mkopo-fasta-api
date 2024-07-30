package models

import (
	"time"
)

type LoanApplication struct {
	ID              uint `gorm:"primaryKey"`
	CustomerID      uint
	Customer        Customer
	LoanAmount      float64   `gorm:"not null"`
	LoanCurrency    string    `gorm:"not null"`
	InterestRate    float64   `gorm:"not null"`
	ApplicationDate time.Time `gorm:"not null"`
	Status          string    `gorm:"type:enum('Pending', 'Approved', 'Rejected');not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
