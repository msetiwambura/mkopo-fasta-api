package models

import (
	"time"
)

type Loan struct {
	ID             uint          `gorm:"primaryKey"`
	CustomerID     uint          `json:"Customer"` // Ensure CustomerID is json tag
	Customer       Customer      `gorm:"foreignKey:CustomerID"`
	LoanAmount     float64       `gorm:"not null"`
	TotalRepayment float64       `gorm:"not null"`
	LoanCurrency   string        `gorm:"not null"`
	InterestRate   float64       `gorm:"not null"`
	StartDate      string        `json:"StartDate" gorm:"type:date;not null"`
	EndDate        string        `json:"EndDate" gorm:"type:date;not null"`
	Status         string        `gorm:"type:enum('Pending', 'Approved', 'Rejected', 'Closed');not null"`
	CreatedAt      time.Time     `gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `gorm:"autoUpdateTime"`
	Payments       []Payment     `gorm:"foreignKey:LoanID"`
	Transactions   []Transaction `gorm:"foreignKey:LoanID"`
	Collaterals    []Collateral  `gorm:"foreignKey:LoanID"`
	Guarantors     []Guarantor   `gorm:"foreignKey:LoanID"`
}

type LoanRes struct {
	ID           uint `gorm:"primaryKey"`
	CustomerID   uint
	CustomerRes  CustomerRes `json:"Customer"`
	LoanAmount   float64     `gorm:"not null"`
	LoanCurrency string      `gorm:"not null"`
	InterestRate float64     `gorm:"not null"`
	StartDate    string      `json:"StartDate" gorm:"type:date;not null"`
	EndDate      string      `json:"EndDate" gorm:"type:date;not null"`
	Status       string      `gorm:"type:enum('Pending', 'Approved', 'Rejected', 'Closed');not null"`
}
