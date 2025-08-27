package dto

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

func NewMovieResponse(id, name string) MovieResponse {
	return MovieResponse{
		ID:   id,
		Name: name,
	}
}
