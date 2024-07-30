package controllers

import (
	"github.com/gin-gonic/gin"
	"loanapi/configs"
	"loanapi/models"
	"loanapi/responses"
	"net/http"
)

func CreateGuarantor(g *gin.Context) {
	var guarantor models.Guarantor
	if err := g.ShouldBindJSON(&guarantor); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the Guarantor
	if err := configs.DB.Create(&guarantor).Error; err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Preload the Loan object
	if err := configs.DB.Preload("Loan").First(&guarantor, guarantor.ID).Error; err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Construct custom responses
	response := responses.CreateSuccessPostResponse("Guarantor Created Successfully", "Ok", guarantor.ID)

	g.JSON(http.StatusOK, response)
}
