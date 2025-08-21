package users

import (
	"errors"
	"net/http"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/creating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/gin-gonic/gin"
)

// CreateUserHandler returns a handler function that processes user creation requests.
func CreateUserHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.UserCreateRequest
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewUserCommand(dto))

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidUserID),
				errors.Is(err, domain.ErrInvalidUserName),
				errors.Is(err, domain.ErrInvalidUserEmail):
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, domain.ErrUserAlreadyExists):
				ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
