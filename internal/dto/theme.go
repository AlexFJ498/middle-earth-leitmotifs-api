package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type ThemeCreateRequest struct {
	Name            string  `json:"name" binding:"required"`
	FirstHeard      string  `json:"first_heard" binding:"required"`
	GroupID         string  `json:"group_id" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	FirstHeardStart int     `json:"first_heard_start" binding:"gte=0"`
	FirstHeardEnd   int     `json:"first_heard_end" binding:"gte=0"`
	CategoryID      *string `json:"category_id"`
}

type ThemeUpdateRequest struct {
	Name            string  `json:"name" binding:"required"`
	FirstHeard      string  `json:"first_heard" binding:"required"`
	GroupID         string  `json:"group_id" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	FirstHeardStart int     `json:"first_heard_start" binding:"gte=0"`
	FirstHeardEnd   int     `json:"first_heard_end" binding:"gte=0"`
	CategoryID      *string `json:"category_id"`
}

type ThemeResponse struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	FirstHeard      TrackResponse     `json:"first_heard"`
	Group           GroupResponse     `json:"group"`
	Description     string            `json:"description"`
	FirstHeardStart int               `json:"first_heard_start"`
	FirstHeardEnd   int               `json:"first_heard_end"`
	Category        *CategoryResponse `json:"category"`
}

func NewThemeResponse(theme domain.Theme, firstHeard TrackResponse, group GroupResponse, category *CategoryResponse) ThemeResponse {
	return ThemeResponse{
		ID:              theme.ID().String(),
		Name:            theme.Name().String(),
		FirstHeard:      firstHeard,
		Group:           group,
		Description:     theme.Description().String(),
		FirstHeardStart: theme.FirstHeardStart().Int(),
		FirstHeardEnd:   theme.FirstHeardEnd().Int(),
		Category:        category,
	}
}
