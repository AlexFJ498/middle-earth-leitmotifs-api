package themes

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/creating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// CreateHandler returns a handler for creating themes.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.ThemeCreateRequest
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewThemeCommand(dto))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidThemeID),
				errors.Is(err, domain.ErrInvalidThemeName),
				errors.Is(err, domain.ErrInvalidGroupID),
				errors.Is(err, domain.ErrInvalidCategoryID),
				errors.Is(err, domain.ErrInvalidTrackID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrGroupNotFound),
				errors.Is(err, domain.ErrCategoryNotFound),
				errors.Is(err, domain.ErrTrackNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
