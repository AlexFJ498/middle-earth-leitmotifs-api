package jwt

import (
	"net/http"
	"strings"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/gin-gonic/gin"
)

// Middleware is a gin.HandlerFunc that middleware for handling JWT authentication.
func Middleware(jwtKey auth.JWTKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate and parse the JWT token
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Expects Bearer <token> format
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			return
		}
		tokenString = parts[1]

		claims, err := auth.ValidateToken(tokenString, jwtKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Store the claims in the context
		c.Set("userID", claims["id"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
