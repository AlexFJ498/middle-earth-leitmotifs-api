package themes

import (
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

// GetHandler handles the getting of a theme.
func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		themeIDParam := ctx.Param("id")
		if themeIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "theme ID is required"})
			return
		}
		theme, err := queryBus.Ask(ctx, getting.NewThemesQuery(themeIDParam))
		if err != nil {
			switch err {
			case domain.ErrGroupNotFound,
				domain.ErrCategoryNotFound,
				domain.ErrTrackNotFound,
				domain.ErrThemeNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case domain.ErrInvalidThemeID:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, theme)
	}
}
