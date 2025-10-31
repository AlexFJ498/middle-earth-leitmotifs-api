package categories

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
		var req dto.CategoryUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cmd := updating.NewCategoryCommand(ctx.Param("id"), req)
		if err := commandBus.Dispatch(ctx, cmd); err != nil {
			switch {
			case errors.Is(err, domain.ErrCategoryNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrInvalidCategoryID),
				errors.Is(err, domain.ErrInvalidCategoryName):
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
