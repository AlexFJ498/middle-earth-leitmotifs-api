package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type TrackCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	MovieID string `json:"movie_id" binding:"required"`
}

type TrackUpdateRequest struct {
	Name    string `json:"name" binding:"required"`
	MovieID string `json:"movie_id" binding:"required"`
}

type TrackResponse struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Movie MovieResponse `json:"movie"`
}

func NewTrackResponse(track domain.Track, movie MovieResponse) TrackResponse {
	return TrackResponse{
		ID:    track.ID().String(),
		Name:  track.Name().String(),
		Movie: movie,
	}
}
