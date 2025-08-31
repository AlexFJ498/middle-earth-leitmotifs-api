package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type ThemeCreateRequest struct {
	Name       string  `json:"name" binding:"required"`
	FirstHeard string  `json:"first_heard" binding:"required"`
	GroupID    string  `json:"group_id" binding:"required"`
	CategoryID *string `json:"category_id"`
}

type ThemeUpdateRequest struct {
	Name       string  `json:"name" binding:"required"`
	FirstHeard string  `json:"first_heard" binding:"required"`
	GroupID    string  `json:"group_id" binding:"required"`
	CategoryID *string `json:"category_id"`
}

type ThemeResponse struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	FirstHeard TrackResponse     `json:"first_heard"`
	GroupID    GroupResponse     `json:"group_id"`
	CategoryID *CategoryResponse `json:"category_id"`
}

func NewThemeResponse(theme domain.Theme, firstHeard TrackResponse, groupID GroupResponse, categoryID *CategoryResponse) ThemeResponse {
	return ThemeResponse{
		ID:         theme.ID().String(),
		Name:       theme.Name().String(),
		FirstHeard: firstHeard,
		GroupID:    groupID,
		CategoryID: categoryID,
	}
}
