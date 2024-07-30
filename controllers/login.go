package controllers

import (
	"github.com/gin-gonic/gin"
	"loanapi/utils"
	"net/http"
)

type LoginRequest struct {
	ChannelID string `json:"ChannelID" binding:"required"`
	IPAddress string `json:"IPAddress" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the ChannelID and IPAddress here (e.g., check against a database)
	// For demonstration, assume they are valid

	token, err := utils.GenerateJWT(req.ChannelID, req.IPAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
