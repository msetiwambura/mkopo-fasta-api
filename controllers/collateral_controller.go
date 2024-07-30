package controllers

import (
	"github.com/gin-gonic/gin"
	"loanapi/configs"
	"loanapi/models"
	"net/http"
)

func CreateCollateral(c *gin.Context) {
	var collateral models.Collateral
	if err := c.ShouldBindJSON(&collateral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configs.DB.Create(&collateral)
	c.JSON(http.StatusCreated, gin.H{"collateral": collateral})
}
