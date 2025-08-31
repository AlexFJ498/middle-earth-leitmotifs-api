package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type MovieCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type MovieUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

type MovieResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewMovieResponse(movie domain.Movie) MovieResponse {
	return MovieResponse{
		ID:   movie.ID().String(),
		Name: movie.Name().String(),
	}
}
