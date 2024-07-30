package models

import "loanapi/configs"

func MigrateDB() {
	configs.DB.AutoMigrate(
		&Customer{},
		&Loan{},
		&Payment{},
		&Employee{},
		&LoanApplication{},
		&Transaction{},
		&Collateral{},
		&Guarantor{},
	)
}
