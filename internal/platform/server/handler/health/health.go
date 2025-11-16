package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	}
}
