package tracks

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
		trackIDParam := ctx.Param("id")
		if trackIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "track ID is required"})
			return
		}

		var req dto.TrackUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		cmd := updating.NewTrackCommand(ctx.Param("id"), req)
		if err := commandBus.Dispatch(ctx, cmd); err != nil {
			switch {
			case errors.Is(err, domain.ErrTrackNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidTrackID),
				errors.Is(err, domain.ErrInvalidTrackName):
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
