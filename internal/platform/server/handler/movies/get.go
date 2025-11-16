package movies

import (
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieIDParam := ctx.Param("id")
		if movieIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "movie ID is required"})
			return
		}
		movie, err := queryBus.Ask(ctx, getting.NewMoviesQuery(movieIDParam))
		if err != nil {
			switch err {
			case domain.ErrMovieNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case domain.ErrInvalidMovieID:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	}
}
