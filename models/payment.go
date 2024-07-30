package models

import (
	"time"
)

type Payment struct {
	ID              uint `gorm:"primaryKey"`
	LoanID          uint
	Loan            Loan
	PaymentAmount   float64   `gorm:"not null"`
	PaymentCurrency string    `gorm:"not null"`
	PaymentDate     string    `json:"PaymentDate" gorm:"type:date;not null"`
	PaymentMethod   string    `gorm:"type:enum('Cash', 'Bank Transfer', 'Credit Card', 'Debit Card');not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
