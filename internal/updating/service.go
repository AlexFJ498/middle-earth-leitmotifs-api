package updating

import (
	"context"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/dto"
)

type MovieService struct {
	movieRepository domain.MovieRepository
}

func NewMovieService(movieRepository domain.MovieRepository) MovieService {
	return MovieService{
		movieRepository: movieRepository,
	}
}

func (s *MovieService) UpdateMovie(ctx context.Context, id string, dto dto.MovieUpdateRequest) error {
	movie, err := domain.NewMovieWithID(id, dto.Name)
	if err != nil {
		return err
	}
	return s.movieRepository.Update(ctx, movie)
}
