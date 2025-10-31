package groups

import (
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupIDParam := ctx.Param("id")
		if groupIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "group ID is required"})
			return
		}
		group, err := queryBus.Ask(ctx, getting.NewGroupsQuery(groupIDParam))
		if err != nil {
			switch err {
			case domain.ErrGroupNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case domain.ErrInvalidGroupID:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, group)
	}
}
