package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type TrackThemeCreateRequest struct {
	TrackID     string `json:"track_id" binding:"required,uuid"`
	ThemeID     string `json:"theme_id" binding:"required,uuid"`
	StartSecond int    `json:"start_second" binding:"required"`
	EndSecond   int    `json:"end_second" binding:"required,gtfield=StartSecond"`
	IsVariant   bool   `json:"is_variant"`
}

type TrackThemeUpdateRequest struct {
	TrackID     string `json:"track_id" binding:"required,uuid"`
	ThemeID     string `json:"theme_id" binding:"required,uuid"`
	StartSecond int    `json:"start_second" binding:"required"`
	EndSecond   int    `json:"end_second" binding:"required,gtfield=StartSecond"`
	IsVariant   bool   `json:"is_variant"`
}

type TrackThemeDeleteRequest struct {
	TrackID     string `json:"track_id" binding:"required,uuid"`
	ThemeID     string `json:"theme_id" binding:"required,uuid"`
	StartSecond int    `json:"start_second" binding:"required"`
}

type TrackThemeResponse struct {
	TrackID     string `json:"track_id"`
	ThemeID     string `json:"theme_id"`
	StartSecond int    `json:"start_second"`
	EndSecond   int    `json:"end_second"`
	IsVariant   bool   `json:"is_variant"`
}

func NewTrackThemeResponse(TrackTheme domain.TrackTheme) TrackThemeResponse {
	return TrackThemeResponse{
		TrackID:     TrackTheme.TrackID().String(),
		ThemeID:     TrackTheme.ThemeID().String(),
		StartSecond: TrackTheme.StartSecond().Int(),
		EndSecond:   TrackTheme.EndSecond().Int(),
		IsVariant:   TrackTheme.IsVariant().Bool(),
	}
}
