package movies

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/creating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// CreateHandler returns a handler function that processes movie creation requests.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.MovieCreateRequest
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewMovieCommand(dto))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidMovieID),
				errors.Is(err, domain.ErrInvalidMovieName):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid movie ID"})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
