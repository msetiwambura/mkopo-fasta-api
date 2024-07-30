package controllers

import (
	"github.com/gin-gonic/gin"
	"loanapi/configs"
	"loanapi/models"
	"loanapi/responses"
	"net/http"
)

func CreatePayment(p *gin.Context) {
	var payment models.Payment
	if err := p.ShouldBindJSON(&payment); err != nil {
		p.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the payment
	if err := configs.DB.Create(&payment).Error; err != nil {
		p.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preload the Loan object
	if err := configs.DB.Preload("Loan").First(&payment, payment.ID).Error; err != nil {
		p.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct custom responses
	response := responses.CreateSuccessPostResponse("Payment Created Successfully", "OK", payment.ID)

	p.JSON(http.StatusOK, response)
}
