package categories

import (
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/getting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryIDParam := ctx.Param("id")
		if categoryIDParam == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "category ID is required"})
			return
		}
		category, err := queryBus.Ask(ctx, getting.NewCategoriesQuery(categoryIDParam))
		if err != nil {
			switch err {
			case domain.ErrCategoryNotFound:
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			case domain.ErrInvalidCategoryID:
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, category)
	}
}
