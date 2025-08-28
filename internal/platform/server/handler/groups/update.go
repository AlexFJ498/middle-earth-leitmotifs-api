package groups

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/updating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// UpdateHandler handles the update of a group.
func UpdateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.GroupUpdateRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cmd := updating.NewGroupCommand(ctx.Param("id"), req)
		if err := commandBus.Dispatch(ctx, cmd); err != nil {
			switch {
			case errors.Is(err, domain.ErrGroupNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			case errors.Is(err, domain.ErrInvalidGroupID),
				errors.Is(err, domain.ErrInvalidGroupName):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		ctx.Status(http.StatusNoContent)
	}
}
