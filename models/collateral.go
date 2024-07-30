package models

import (
	"time"
)

type Collateral struct {
	ID             uint `gorm:"primaryKey"`
	LoanID         uint
	Loan           Loan
	CollateralType string    `gorm:"size:100;not null"`
	Description    string    `gorm:"type:text;not null"`
	Value          float64   `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
