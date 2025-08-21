package session

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/authenticating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler handles user login requests.
func LoginHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := queryBus.Ask(ctx, authenticating.NewLoginQuery(req.Email, req.Password))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidUserPassword):
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidUserEmail):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
