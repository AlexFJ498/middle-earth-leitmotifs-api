package groups

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
		groupIDParam := ctx.Param("id")
		if groupIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "group ID is required"})
			return
		}

		err := commandBus.Dispatch(ctx, deleting.NewGroupCommand(groupIDParam))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidGroupID):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrGroupNotFound):
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
