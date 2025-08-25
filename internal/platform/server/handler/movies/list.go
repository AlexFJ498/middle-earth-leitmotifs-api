package movies

import (
	"net/http"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/listing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

// ListHandler handles the listing of movies.
func ListHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movies, err := queryBus.Ask(ctx, listing.NewMoviesQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movies)
	}
}
