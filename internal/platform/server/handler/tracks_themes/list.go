package tracks_themes

import (
	"net/http"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/listing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func ListByTrackHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackID := ctx.Param("id")
		tracksThemes, err := queryBus.Ask(ctx, listing.NewTracksThemesByTrackQuery(trackID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, tracksThemes)
	}
}
