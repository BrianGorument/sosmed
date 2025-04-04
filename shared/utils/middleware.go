package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is used to check the validity of the JWT token in the Authorization header.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Token should be in the format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token and extract user information
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		
		// Extract the claims for user data
		userID, _ := claims["userId"].(float64)	
		userName, _ := claims["userName"].(string)
		userEmail, _ := claims["userEmail"].(string)
		// Set userID in context for further use in the handler	
		c.Set("userId", userID)
		c.Set("userName", userName)
		c.Set("userEmail", userEmail)
		c.Next()
	}
}
