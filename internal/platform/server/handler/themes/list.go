package themes

import (
	"net/http"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/listing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func ListHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		themes, err := queryBus.Ask(ctx, listing.NewThemesQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, themes)
	}
}

func ListByGroupHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupID := ctx.Param("group_id")
		themes, err := queryBus.Ask(ctx, listing.NewThemesByGroupQuery(groupID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, themes)
	}
}
