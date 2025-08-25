package movies

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/deleting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieIDParam := ctx.Param("id")
		if movieIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "movie ID is required"})
			return
		}

		err := commandBus.Dispatch(ctx, deleting.NewMovieCommand(movieIDParam))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidMovieID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid movie ID"})
				return
			case errors.Is(err, domain.ErrMovieNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.Status(http.StatusNoContent)
	}
}
