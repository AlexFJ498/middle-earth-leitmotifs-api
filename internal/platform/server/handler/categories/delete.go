package categories

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
		categoryIDParam := ctx.Param("id")
		if categoryIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "category ID is required"})
			return
		}

		err := commandBus.Dispatch(ctx, deleting.NewCategoryCommand(categoryIDParam))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidCategoryID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrCategoryNotFound):
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
