package movies

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/updating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieIDParam := ctx.Param("id")
		if movieIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "movie ID is required"})
			return
		}

		var req dto.MovieUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cmd := updating.NewMovieCommand(movieIDParam, req)
		if err := commandBus.Dispatch(ctx, cmd); err != nil {
			switch {
			case errors.Is(err, domain.ErrMovieNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidMovieID),
				errors.Is(err, domain.ErrInvalidMovieName):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.Status(http.StatusNoContent)
	}
}
