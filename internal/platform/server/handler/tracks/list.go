package tracks

import (
	"net/http"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/listing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func ListHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tracks, err := queryBus.Ask(ctx, listing.NewTracksQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, tracks)
	}
}

func ListByMovieHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieID := ctx.Param("id")
		tracks, err := queryBus.Ask(ctx, listing.NewTracksByMovieQuery(movieID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, tracks)
	}
}
