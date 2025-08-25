package deleting

import (
	"context"
	"fmt"
	"testing"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMovieServiceDeleteMovieRepositoryError(t *testing.T) {
	movieIDObj, err := domain.NewMovieIDFromString("123e4567-e89b-12d3-a456-426614174000")
	require.NoError(t, err)

	mockRepo := new(storagemocks.MovieRepository)
	mockRepo.On("Delete", mock.Anything, movieIDObj).Return(fmt.Errorf("database error"))

	service := NewMovieService(mockRepo)

	err = service.DeleteMovie(context.Background(), movieIDObj)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")

	mockRepo.AssertExpectations(t)
}

func TestMovieServiceDeleteMovieSuccess(t *testing.T) {
	movieIDObj, err := domain.NewMovieIDFromString("123e4567-e89b-12d3-a456-426614174000")
	require.NoError(t, err)

	mockRepo := new(storagemocks.MovieRepository)
	mockRepo.On("Delete", mock.Anything, movieIDObj).Return(nil)

	service := NewMovieService(mockRepo)

	err = service.DeleteMovie(context.Background(), movieIDObj)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
