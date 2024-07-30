package models

import (
	"time"
)

type Transaction struct {
	ID              uint `gorm:"primaryKey"`
	LoanID          uint
	Loan            Loan
	CustomerID      uint
	Customer        Customer
	Amount          float64   `gorm:"not null"`
	TransactionType string    `gorm:"type:enum('Disbursement', 'Repayment');not null"`
	TransactionDate time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
