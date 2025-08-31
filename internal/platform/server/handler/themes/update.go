package themes

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/updating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// UpdateHandler returns the handler that updates themes.
func UpdateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		themeIDParam := ctx.Param("id")
		if themeIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "theme ID is required"})
			return
		}

		var req dto.ThemeUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		cmd := updating.NewThemeCommand(ctx.Param("id"), req)
		if err := commandBus.Dispatch(ctx, cmd); err != nil {
			switch {
			case errors.Is(err, domain.ErrThemeNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidThemeID),
				errors.Is(err, domain.ErrInvalidThemeName):
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
