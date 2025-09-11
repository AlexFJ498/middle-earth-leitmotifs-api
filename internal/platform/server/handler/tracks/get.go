package tracks

import (
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

// GetHandler handles the getting of a track.
func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trackIDParam := ctx.Param("id")
		if trackIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "track ID is required"})
			return
		}
		track, err := queryBus.Ask(ctx, getting.NewTracksQuery(trackIDParam))
		if err != nil {
			switch err {
			case domain.ErrMovieNotFound,
				domain.ErrTrackNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case domain.ErrInvalidTrackID:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, track)
	}
}
