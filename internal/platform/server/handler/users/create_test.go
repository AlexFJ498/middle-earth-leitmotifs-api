package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command/commandmocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const usersRoute = "/users"

func TestCreateHandler(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On("Dispatch", mock.Anything, mock.AnythingOfType("creating.UserCommand")).Return(nil)
	defer commandBus.AssertExpectations(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(usersRoute, CreateHandler(commandBus))

	t.Run("Given invalid request, should return 400", func(t *testing.T) {
		createUserReq := createRequest{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		b, err := json.Marshal(createUserReq)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", usersRoute, bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Given valid request, should return 201", func(t *testing.T) {
		createUserReq := createRequest{
			ID:    "550e8400-e29b-41d4-a716-446655440000",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		b, err := json.Marshal(createUserReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, usersRoute, bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
