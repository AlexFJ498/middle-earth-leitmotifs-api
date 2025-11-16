package tracks_themes

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/deleting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.TrackThemeDeleteRequest
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(ctx, deleting.NewTrackThemeCommand(dto))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidTrackID),
				errors.Is(err, domain.ErrInvalidThemeID),
				errors.Is(err, domain.ErrInvalidStartSecond):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrTrackNotFound),
				errors.Is(err, domain.ErrThemeNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.Status(http.StatusNoContent)
	}
}
