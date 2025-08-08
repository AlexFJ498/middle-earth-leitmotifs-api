package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler returns a handler function that responds with a health check status.
func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}
}
