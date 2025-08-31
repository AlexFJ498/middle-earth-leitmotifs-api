package themes

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/deleting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// DeleteHandler returns the handler that deletes themes.
func DeleteHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "theme ID is required"})
			return
		}

		err := commandBus.Dispatch(ctx, deleting.NewThemeCommand(id))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidThemeID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrThemeNotFound):
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
