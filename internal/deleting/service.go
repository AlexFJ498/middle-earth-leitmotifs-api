package deleting

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
)

// MovieService is the default implementation of the MovieService interface.
type MovieService struct {
	movieRepository domain.MovieRepository
}

// NewMovieService returns a new MovieService instance.
func NewMovieService(repo domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: repo,
	}
}

// DeleteMovie deletes a movie by its ID.
func (s *MovieService) DeleteMovie(ctx context.Context, id domain.MovieID) error {
	return s.movieRepository.Delete(ctx, id)
}
