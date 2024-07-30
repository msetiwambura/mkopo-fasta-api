package middlewares

import (
	"github.com/gin-gonic/gin"
	"loanapi/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Check for "Bearer " prefix and strip it
		const prefix = "Bearer "
		if !strings.HasPrefix(authorizationHeader, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "tokenstring should contain 'Bearer ' prefix"})
			c.Abort()
			return
		}

		tokenString := authorizationHeader[len(prefix):] // Strip "Bearer " prefix

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		expectedChannelID := c.GetHeader("ChannelID") // Replace with the actual channel ID you expect
		if claims.ChannelID != expectedChannelID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Channel ID mismatch"})
			c.Abort()
			return
		}

		// Set claims in context for later use
		c.Set("ChannelID", claims.ChannelID)
		c.Set("IPAddress", claims.IPAddress)

		c.Next()
	}
}
